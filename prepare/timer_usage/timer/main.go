package main

import (
	"fmt"
	"time"
)

func main()  {
	input := make(chan interface{})

	go func() {
		for i:=0;i<5 ;i++  {
			input <- i
		}
		input <- "hello word"
	}()

	t1 := time.NewTimer(time.Second * 5)
	t2 := time.NewTimer(time.Second * 10)

	for  {
		select {
		case msg := <- input:
			fmt.Println(msg)
		case <- t1.C:
			fmt.Println("5s")
			t1.Reset(time.Second * 5)
		case <-t2.C:
			fmt.Println("10s")
			t2.Reset(time.Second * 10)
			
		}
	}
}
