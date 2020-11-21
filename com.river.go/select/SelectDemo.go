package main

import (
	"fmt"
	"time"
)

func main() {

	var a = make(chan int)
	var b = make(chan int)

	go func() {
		b <- 22
		a <- 11

	}()
	<-time.After(time.Second * 10)
	//fmt.Println(ii)
	time.Sleep(time.Second)
	for i := 0; i < 2; i++ {
		select {
		case i := <-a:
			fmt.Println(i)
		case j := <-b:
			fmt.Println(j)
		default:
			fmt.Println("null")
		}
	}

}
