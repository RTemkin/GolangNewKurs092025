package main

import (
	"fmt"
	"sync"
)

//var number int = 0

// создание атомарного значения
//var number atomic.Int64

var slice []int

//мютекс
var mtx sync.Mutex


func increace(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		//number++

		//для работы с атомиками используються методы в пакете atomic
		//number.Add(1)
		mtx.Lock()
		slice = append(slice, i)
		mtx.Unlock()

	}
}

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increace(wg)
	}

	wg.Wait()
	// вывод атомика методом .Load()
	//fmt.Println(number.Load())

	fmt.Println(len(slice))

}
