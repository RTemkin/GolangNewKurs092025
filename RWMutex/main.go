package main

import (
	"fmt"
	"sync"
	"time"
)

var likes int = 0
var mtx sync.RWMutex

func getLike(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 100000; i++ {
		mtx.Lock()
		likes++
		mtx.Unlock()
	}
}

func setLike(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		//операция чтения
		mtx.RLock()
		_ = likes
		mtx.RUnlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}

	initTime := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go getLike(wg)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go setLike(wg)
	}

	wg.Wait()
	fmt.Println("Время выполнения:", time.Since(initTime))

}
