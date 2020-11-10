package main

import "fmt"

func main() {
	//数组循环
	array := [5]string{"a", "b", "c", "d", "e"}
	for i := 0; i < 5; i++ {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, array[i])
	}
	for i, v := range array {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
	}
	//切片
	slice := array[2:5]
	fmt.Println(slice)

	slice1 := make([]string, 4, 4)

	strings := append(slice1, "1")
	strings[0] = "666"
	fmt.Println(slice1)

	fmt.Println("===================>")
	nameAgeMap := make(map[string]int)
	nameAgeMap["张飞"] = 30
	nameAgeMap["哪吒"] = 10
	nameAgeMap["river"] = 33
	nameAgeMap["frank"] = 90
	fmt.Println(nameAgeMap)
	age, ok := nameAgeMap["张飞"]
	fmt.Println(age, ok)
	delete(nameAgeMap, "river")

	for k, v := range nameAgeMap {
		fmt.Println(k, v)
	}

	s := "Hello飞雪无情"
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Println(s[0], s[1], s[15])

	fmt.Println("===================>")
	for i, v := range s {
		fmt.Println(i, v)
	}

}
