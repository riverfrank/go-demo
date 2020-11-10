package main

import (
	"fmt"
	"strings"
)

func main() {
	i := strings.Index("飞雪无情", "飞雪")
	fmt.Println(i)

	if i > 0 {
		fmt.Println("hello")
	} else {
		fmt.Println("no")
	}

	switch j := 1; {
	case j == 1:
		fmt.Println("1")
		fallthrough
	case j == 2:
		fmt.Println("2")
	case j == 3:
		fmt.Println("2")
	}

	sum := 0
	for i := 1; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println(sum)

	fmt.Println("hello river")

}
