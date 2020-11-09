package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var i int = 10
	var b bool = false

	pi := &i
	formatBool  := strconv.FormatBool(b)
	fmt.Println("hello river " ,i)
	fmt.Println(b)
	fmt.Println(formatBool)
	fmt.Println(*pi)
	fmt.Println(pi)

	*pi = 99
	fmt.Println(i)

	contains := strings.Contains("hello river", "river")
	fmt.Println(contains)
}
