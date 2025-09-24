package main

import (
	"githab/rtemkin/golangnewkurs092025/interfaces/payments"
	"githab/rtemkin/golangnewkurs092025/interfaces/payments/metods"

	"github.com/k0kubun/pp"
)

func main() {

	metod := metods.NevCripto()

	paymentModul := payments.NewPaymentModul(metod)

	paymentModul.Pay("Burg", 5)
	paymentModul.Pay("telefon", 500)
	paymentModul.Pay("Game", 50)

	allinfo := paymentModul.AllInfo()

	pp.Println("All", allinfo)

}
