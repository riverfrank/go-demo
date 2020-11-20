package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	 * 测试结果
	 * channel无缓冲时，发送阻塞直到数据被接收，接收阻塞直到读到数据。
	 * channel有缓冲时，当缓冲满时发送阻塞，当缓冲空时接收阻塞。
	 */
	ch := make(chan int, 0)
	go func() {
		var sum int
		for i := 0; i < 10; i++ {
			sum += i
		}
		fmt.Println(sum)
		ch <- sum
		//ch <- 22
		fmt.Println("--------->")

	}()
	time.Sleep(time.Second * 2)
	fmt.Println("hello <- ch", <-ch)
	//fmt.Println("hello <- ch", <-ch)
	time.Sleep(time.Second * 1)

}
