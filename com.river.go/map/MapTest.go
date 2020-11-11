package main

import "fmt"

func main() {
	nameMap := make(map[string]int)
	nameMap["张飞"] = 40
	nameMap["赵云"] = 30
	nameMap["river"] = 33

	fmt.Println(nameMap)
	delete(nameMap, "cat")
	delete(nameMap, "赵云")
	age := nameMap["river"]
	fmt.Println(age)
	fmt.Println(nameMap)

	for s, i := range nameMap {
		fmt.Println(s, i)
	}

}
