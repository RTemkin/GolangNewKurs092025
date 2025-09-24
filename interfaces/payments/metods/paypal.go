package metods

import (
	"fmt"
	"math/rand"
)

type Paypal struct{}

func NevPaypal() Paypal {
	return Paypal{}
}

func (p Paypal) Pay(usd int) int {
	fmt.Println("Оплата Paypal")
	fmt.Println("Размер оплаты", usd, "USD")

	id := rand.Int()
	return id

}

func (p Paypal) Cansel(id int) {
	fmt.Println("Отмена Paypal-операции, ID: ", id)
}
