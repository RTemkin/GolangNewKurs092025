package main

import (
	"fmt"
	"sync"
)

func rangeChan(ch <-chan int, mergeCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		mergeCh <- val
	}
}

func mergeC(chans ...chan int) chan int {
	wg := &sync.WaitGroup{}
	merChan := make(chan int)

	for _, ch := range chans {
		wg.Add(1)
		go rangeChan(ch, merChan, wg)
	}

	go func() {
		wg.Wait()
		close(merChan)
	}()

	return merChan
}

func fibonachi(n int) int {
	if n == 0 {
		return 0
	}
	a := 0
	b := 1

	for i := 0; i < n-1; i++ {
		a, b = b, a+b
	}

	return b
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	sliseVal := []int{}

	go func() {
		for i := 0; i <= 30; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 31; i <= 70; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	go func() {
		for i := 71; i <= 100; i++ {
			ch3 <- i
		}
		close(ch3)
	}()

	merCh := mergeC(ch1, ch2, ch3)

	for val := range merCh {
		sliseVal = append(sliseVal, val)
	}

	fmt.Println(sliseVal, len(sliseVal))

	for i := 0; i < 10; i++ {

		fmt.Println(fibonachi(i))
	}
}
