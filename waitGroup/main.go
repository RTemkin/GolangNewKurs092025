package main

import (
	"fmt"
	"sync"
	"time"
)

func postman(text string, wg *sync.WaitGroup) {

	defer wg.Done()
	
	for i := 1; i <= 3; i++ {
		fmt.Println("Отнес", text, "postman-", i)
		time.Sleep(250 * time.Millisecond)
	}

}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go postman("Tims", wg)

	wg.Add(1)
	go postman("Game", wg)

	wg.Add(1)
	go postman("Auto", wg)

	wg.Wait()

	fmt.Println("main завершен!")
}
