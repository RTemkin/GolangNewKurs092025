package metods

import (
	"fmt"
	"math/rand"
)

type Bank struct{}

func NevBank() Bank {
	return Bank{}
}

func (b Bank) Pay(usd int) int {
	fmt.Println("Оплата банк")
	fmt.Println("Размер оплаты", usd, "USD")

	id := rand.Int()
	return id

}

func (b Bank) Cansel(id int) {
	fmt.Println("Отмена Банк-операции, ID: ", id)
}
