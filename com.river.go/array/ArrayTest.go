package main

import "fmt"

func main() {
	//一维数组
	//var myArray = [5]string{"a", "b", "c", "d", "e"}
	myArray := [5]string{"a", "b", "c", "d", "e"}

	fmt.Println(myArray)
	//二维数组
	var my2Array = [2][2]string{
		{"river", "frank"},
		{"bill", "lucy"},
	}
	fmt.Println(my2Array)

	slice := myArray[0:2]
	fmt.Println(slice)
}
