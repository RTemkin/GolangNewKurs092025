package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Payment struct {
	Description string `json:"description"`
	USD         int    `json:"usd"`
	FullName    string `json:"fullname"`
	Address     string `json:"address"`
	Time        time.Time
}

type HttpResponse struct {
	Money          int       `json:"money"`
	PaymentHistory []Payment `json:"paymenhistory"`
}

var mu = sync.Mutex{}
var money = 1000
var paymentHistory = make([]Payment, 0)

func payHendler(w http.ResponseWriter, r *http.Request) {
	var payment Payment

	//более простой способ декодирования тела запроса json
	// err2 := json.NewDecoder(r.Body).Decode(&payment)
	// if err2 != nil {
	// 	fmt.Println("err1:", err2)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	httpRequestBady, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err1 := json.Unmarshal(httpRequestBady, &payment)
	if err1 != nil {
		fmt.Println("err1:", err1)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	payment.Time = time.Now()

	mu.Lock()
	if money >= payment.USD {
		money -= payment.USD
	}

	paymentHistory = append(paymentHistory, payment)

	httpResponse := HttpResponse{
		Money:          money,
		PaymentHistory: paymentHistory,
	}

	b, err := json.Marshal(httpResponse)
	if err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("Money: ", money)
	fmt.Println("PaymentHistory: ", paymentHistory)
	mu.Unlock()
}


// Qvery параметры например youtube.com/video?v=123&t=361

//localhost:9091/default?foo=x&boo=y
func defaultHendler(w http.ResponseWriter, r *http.Request){
	fooParam := r.URL.Query().Get("foo")
	booParam := r.URL.Query().Get("boo")

	fmt.Println("fooParam:", fooParam)
	fmt.Println("booParam:", booParam)
}


func main() {
	http.HandleFunc("/pay", payHendler)

	http.HandleFunc("/default", defaultHendler)

	fmt.Println("Запуск сервера")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка во время запуска http сервера", err)
	}

}
