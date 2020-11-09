package main

import "fmt"

func main1() {
	var a int = 0
	var b int
	c := 3

	var n [10]int /* n 是一个长度为 10 的数组 */


	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var i,j int
	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j] )
	}

	var aa = [3][4]int{
		{0, 1, 2, 3} ,   /*  第一行索引为 0 */
		{4, 5, 6, 7} ,   /*  第二行索引为 1 */
		{8, 9, 10, 11},   /* 第三行索引为 2 */
	}
	fmt.Println(aa[1][1])

	fmt.Println("----range test----")
	nums := [] int {2,3,4}
	//sum := 0
	for i,num := range nums {
		fmt.Println(i,num)
	}
}
