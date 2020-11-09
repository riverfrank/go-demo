package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}
func max(num1,num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
func main2() {

	a , b := swap("hello","river")
	fmt.Println(a,b)
	fmt.Println(max(2,9))


}
