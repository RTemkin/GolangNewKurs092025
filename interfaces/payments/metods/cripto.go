package metods

import (
	"math/rand"
	"fmt"
)

type Cripto struct{}

func NevCripto()Cripto{
	return Cripto{}
}

func (c Cripto) Pay(usd int) int {
	fmt.Println("Оплата крипнтовалютой")
	fmt.Println("Размер оплаты", usd, "USDT")

	id := rand.Int()
	return id

}

func (c Cripto)Cansel(id int){
	fmt.Println("Отмена крипто-операции, ID: ", id )
}
