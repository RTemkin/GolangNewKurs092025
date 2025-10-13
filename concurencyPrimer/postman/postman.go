package postman

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Postman(ctx context.Context, transferPoint chan<- string, n int, mail string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Почтальон %v завершил работу!!!!!\n", n)
			return
		default:
			fmt.Printf("Почтальон %v взял письмо\n", n)
			time.Sleep(time.Second)
			fmt.Printf("%v - почтальем отнес письмо - %v\n", n, mail)

			transferPoint <- mail

			fmt.Printf("%v - почтальем передол письмо!!! письмо - %v\n", n, mail)
		}
	}
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)

	wg := &sync.WaitGroup{}

	for i := 1; i <= postmanCount; i++ {
		mail := postmanToMail(i)
		wg.Add(1)
		go Postman(ctx, mailTransferPoint, i, mail, wg)
	}

	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint
}

func postmanToMail(postmanNumber int) string {
	ptm := map[int]string{
		1: "aaaaaaaaaaaaaaa",
		2: "bbbbbbbbbbbbbbb",
		3: "ccccccccccccccc",
	}

	mail, ok := ptm[postmanNumber]
	if !ok {
		mail = ptm[rand.Intn(3)+1]
	}

	return mail
}
