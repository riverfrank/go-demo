package main

import (
	"fmt"
	"go-demo/com.river.go/flow"
	"strconv"
	"strings"
)

func main() {

	flow.SayHello()

	var i int = 10
	var b bool = false
	var x, y int
	fmt.Println(x, y)

	pi := &i
	formatBool := strconv.FormatBool(b)
	fmt.Println("hello river ", i)
	fmt.Println(b)
	fmt.Println(formatBool)
	fmt.Println(*pi)
	fmt.Println(pi)

	*pi = 99
	fmt.Println(i)

	contains := strings.Contains("hello river", "river")
	fmt.Println(contains)

}
