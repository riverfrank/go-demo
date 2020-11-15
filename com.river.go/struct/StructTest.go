package main

import "fmt"

type person struct {
	name string
	age  int
}

func modifyAge(p *person) {
	p.age = 33
}
func main() {
	lucy := person{name: "lucy", age: 22}
	var lucyP *person = &lucy
	fmt.Println(lucy)
	modifyAge(&lucy)
	fmt.Println(lucy)
	/**
	 * c 语言的写法 (*person).name
	 * go 支持两种写法 (*lucyP).name and lucyP.name
	 * 可见 go 默认做了推导
	 */
	fmt.Println((*lucyP).name)
	fmt.Println(lucyP.name)

}
