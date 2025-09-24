package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/k0kubun/pp"
)

type User struct {
	Name     string
	Ballance int
}

func Pay(user *User, usd int) error {
	if user.Ballance-usd < 0 {
		err := errors.New("недостаточно средств")
		return err
	}

	user.Ballance -= usd
	return nil
}

func main() {

	user := User{
		Name:     "Oleg",
		Ballance: 100,
	}

	pp.Println(user)
	err := Pay(&user, 34)

	pp.Println(user)

	if err != nil {
		fmt.Println("Оплаты нет, ПРичина:", err.Error())
	} else {
		fmt.Println("Оплата прошла!")
	}

	fmt.Println("---------------------------------------")
	//panic
	// лучше в самом начале функции в которой возможна паника
	defer func() {
		p := recover()
		if p != nil {
			log.Println("Panic recover", p)
		}
	}()

	a := 0
	b := 1 / a

	fmt.Println(b)

	fmt.Println("After panic func recover")

}
