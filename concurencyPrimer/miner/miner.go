package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Miner(ctx context.Context, transferPoint chan<- int, n int, power int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Шахтер %v закончил работу!!!\n", n)
			return
		default:
			fmt.Printf("Шахтер %v начал добывать уголь\n", n)
			time.Sleep(time.Second)
			fmt.Printf("%v - шахтер добыл уголь\n", n)

			transferPoint <- power

			fmt.Printf("%v - шахтер передал %v уголz\n", n, power)
		}
	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)

	wg := &sync.WaitGroup{}

	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go Miner(ctx, coalTransferPoint, i, i*10, wg)
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}
