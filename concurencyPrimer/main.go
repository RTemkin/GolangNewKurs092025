package main

import (
	"context"
	"fmt"
	"githab/rtemkin/golangnewkurs092025/concurencyPrimer/miner"
	"githab/rtemkin/golangnewkurs092025/concurencyPrimer/postman"
	"sync"
	"time"
)

func main() {

	var coal int
	var mails []string

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	coalTransferPoint := miner.MinerPool(minerContext, 10)
	mailTransferPoint := postman.PostmanPool(postmanContext, 10)

	go func() {
		time.Sleep(2 * time.Second)
		minerCancel()
	}()

	go func() {
		time.Sleep(3 * time.Second)
		postmanCancel()
	}()

	wg := &sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range coalTransferPoint {
			mu.Lock()
			coal += val
			mu.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range mailTransferPoint {
			mu.Lock()
			mails = append(mails, val)
			mu.Unlock()
		}
	}()

	// isCoalClosed := false
	// isMailClosed := false
	// for !isCoalClosed || !isMailClosed {
	// 	select {
	// 	case c, ok := <-coalTransferPoint:
	// 		if !ok {
	// 			isCoalClosed = true
	// 			continue
	// 		}
	// 		coal += c
	// 	case m, ok := <-mailTransferPoint:
	// 		if !ok {
	// 			isMailClosed = true
	// 			continue
	// 		}
	// 		mails = append(mails, m)
	// 	}
	// }

	wg.Wait()

	mu.Lock()
	fmt.Println("Coal:", coal)
	mu.Unlock()

	mu.Lock()
	fmt.Println("malis:", mails)
	mu.Unlock()

}
