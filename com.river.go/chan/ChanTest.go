package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		var sum int
		for i := 0; i < 10; i++ {
			sum += i
		}
		fmt.Println(sum)
		ch <- sum

		ch <- 5
		ch <- 4
		ch <- 3
		ch <- 2
		fmt.Println(2)

	}()
	fmt.Println("hello <- ch", <-ch)
	fmt.Println("hello <- ch", <-ch)
}
