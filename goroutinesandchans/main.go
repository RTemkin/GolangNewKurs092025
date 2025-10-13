package main

import (
	"fmt"
	"time"
)

func mine(ch1 chan int, n int) {
	fmt.Printf("Начало похода %v\n", n)
	time.Sleep(1 * time.Second)
	fmt.Printf("Конец похода %v\n", n)

	ch1 <- 10
	fmt.Println("Передал", n)
}

func main() {
	coal := 0
	transferChan := make(chan int, 10)

	initTime := time.Now()
	
	go mine(transferChan, 1)
	go mine(transferChan, 2)
	go mine(transferChan, 3)
	go mine(transferChan, 4)
	go mine(transferChan, 5)
	go mine(transferChan, 6)
	go mine(transferChan, 7)
	go mine(transferChan, 8)
	go mine(transferChan, 9)

	fmt.Println(time.Since(initTime))

	for i := 0; i < 9; i++ {
		coal += <-transferChan
	}
	close(transferChan)

	fmt.Println(coal)
	fmt.Println("Время выполнения:", time.Since(initTime))

}
