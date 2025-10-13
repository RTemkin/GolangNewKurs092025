package main

import (
	"fmt"
	"sync"
)

func squar(n int, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- n*n

}

func sumSquar(n int) int {
	wg := &sync.WaitGroup{}
	out := make(chan int, n)
	sum := 0
	for i := 1; i<=n; i++{
		wg.Add(1)
		go squar(i, out, wg)
	}

	wg.Wait()
	close(out)

	for val := range out{
		sum +=val
	}

	return sum
}

func main() {

	
	fmt.Println(sumSquar(5))

}
