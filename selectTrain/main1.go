package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	intCh := make(chan int)
	strCh := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		intCh <- 10
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		strCh <- "hi" + strconv.Itoa(1)
	}()

	time.Sleep(500 * time.Millisecond)

	select {
	case number := <-intCh:
		fmt.Println("intCh", number)
	case str := <-strCh:
		fmt.Println("strCh", str)
	default:
		fmt.Println("Ни какой канал не готов")

	}
}
