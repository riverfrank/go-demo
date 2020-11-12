package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("hello river")
	fmt.Println("max num is ", max(10, 33))

	x, y := swap(10, 33)
	fmt.Println("swat num is ", x, y)

	m, n := sum(11, 1)

	fmt.Println(m, n)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func swap(a, b int) (int, int) {
	return b, a
}
func sum(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("不能是负数")
	}
	return a + b, nil
}
