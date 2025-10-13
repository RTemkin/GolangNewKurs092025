package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var money = 1000
var bank = 0
var mu sync.Mutex

func payHandler(w http.ResponseWriter, r *http.Request) {

	// for key, val := range r.Header {
	// 	// в http header (r.Header) заинта служебная информауция
	// 	fmt.Printf("key: %v, value: %v\n", key, val)
	// }

	// http методы
	fmt.Println("http method:", r.Method)

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		msg := "fatal to reader http body" + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fatal to write http responcse", err)
		}
		return
	}

	httpRequestBodySring := string(httpRequestBody)
	paymentAmount, err := strconv.Atoi(httpRequestBodySring)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "fail to convert http body to int" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fatal to write http responcse-2", err)
		}
		return
	}

	//необходимо брать блокировку на полностью блок операций, а не на только переменную,
	// так как второй handler может так же проводить проверку и в конце уйдем в минус по отсткам денег
	mu.Lock()
	if money >= paymentAmount {
		//mu.Lock()
		money -= paymentAmount
		//mu.Unlock()
		msg := "Оплата прошла успешно! Money:" + strconv.Itoa(money)
		fmt.Println("Оплата прошла успешно! Money:", money)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fatal to write http responcse-3", err)
		}
	} else {
		msg := "Not eguout money, money:" + strconv.Itoa(money)
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fatal to write http responcse-3-1", err)
		}
	}
	mu.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "fail error to request body" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fatal to write http responcse-4", err)
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		msg := "fail error to convert body to int" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fatal to write http responcse-5", err)
		}
		return
	}

	mu.Lock()
	if money >= saveAmount {
		//mu.Lock()
		money -= saveAmount

		bank += saveAmount
		//mu.Unlock()

		msg := "new money:" + strconv.Itoa(money) + " new bank:" + strconv.Itoa(bank)
		fmt.Printf(msg)
		_, err := w.Write([]byte(msg))
		fmt.Println("fatal to write http rensponse - 6", err)
	} else {
		msg := "Not eguout money, money:" + strconv.Itoa(money)
		fmt.Printf(msg)
		_, err := w.Write([]byte(msg))
		fmt.Println("fatal to write http rensponse - 7", err)
	}
	mu.Unlock()
}

func main() {

	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	fmt.Println("Запускаю сервер")

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("произошла ошибка HTTP serv pae-save", err.Error())
	}

	fmt.Println("Программа завершена")
}
