package main

import (
	"fmt"
	"net/http"
)

//localhost == (ip 127.0.0.1)
//postman

func payHandler(w http.ResponseWriter, r *http.Request) {
	str := "Pay"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("произошла ошибка во время запроса http", err)
	} else {
		fmt.Println("Корректно оплачено")
	}
}

func canselhandler(w http.ResponseWriter, r *http.Request) {
	str := "Cancel"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("произошла ошибка во время запроса http", err)
	} else {
		fmt.Println("Корректная отмена")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	str := "Hello World"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("произошла ошибка во время запроса http", err)
	} else {
		fmt.Println("Корректно обработаный http запрос")
	}
}

func main() {
	http.HandleFunc("/default", handler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", canselhandler)

	fmt.Println("Запускаю сервер")

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("произошла ошибка serv", err)
	}

	fmt.Println("Программа завершена")
}
