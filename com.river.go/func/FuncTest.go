package main

import "fmt"

func main() {
	fmt.Println("hello river")
	fmt.Println("max num is ", max(10, 33))

	x, y := swap(10, 33)
	fmt.Println("swat num is ", x, y)
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
