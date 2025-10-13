package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func foo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("foo завершен")
			return
		default:
			fmt.Println("foo продолжает выполнение")
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func boo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("boo завершен")
			return
		default:
			fmt.Println("boo продолжает выполнение")
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int)
	go func() {
		iteration := 3 + rand.Intn(5)
		for i := 0; i < iteration; i++ {
			ch <- 10

		}
		close(ch)
	}()

	coal := 0
	for v := range ch {
		coal += v
	}
	fmt.Println(coal)

	fmt.Println("----------------------------------------")

	//context - контекст

	parentContext, parentCansel := context.WithCancel(context.Background())
	childContext, childCansel := context.WithCancel(parentContext)

	go foo(parentContext)
	go boo(childContext)

	//отменяем childcontext - прекращает работу boo через секунду
	time.Sleep(1 * time.Second)
	childCansel()

	//отменяем parentcontext - прекращает работу foo через секунду
	time.Sleep(1 * time.Second)
	parentCansel()

}
