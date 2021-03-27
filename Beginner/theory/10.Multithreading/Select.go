package main

import (
	"fmt"
	"time"
)

func Select() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "form 1"
			time.Sleep(time.Second * 1)
		}
	}()
	go func() {
		for {
			c2 <- "form 2"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			//select{
			//case msg:=<-c1:
			//	fmt.Println(msg)
			//case msg2:=<-c2:
			//	fmt.Println(msg2)
			//}
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2", msg2)
			case msg3 := <-time.After(time.Second):
				fmt.Println("timeout", msg3)
			default:
				fmt.Println("nothing ready")
				time.Sleep(time.Millisecond * 100)
			}
		}

	}()
	var input string
	fmt.Scanln(&input)
}
