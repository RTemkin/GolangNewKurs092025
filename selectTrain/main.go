package main

import (
	"fmt"
	"time"
)

type Message struct {
	Author string
	Text   string
}

func main() {
	messageChan1 := make(chan Message)
	messageChan2 := make(chan Message)

	go func() {
		for {
			messageChan1 <- Message{
				Author: "Friend1",
				Text:   "Hello",
			}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			messageChan2 <- Message{
				Author: "Friend2",
				Text:   "how are you",
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		select {
		case msg1 := <-messageChan1:
			fmt.Println("Получено от ", msg1.Author, msg1.Text)
		case msg2 := <-messageChan2:
			fmt.Println("Получено от ", msg2.Author, msg2.Text)
		}
	}
}
