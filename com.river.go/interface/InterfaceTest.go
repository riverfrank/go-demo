package main

import "fmt"

type animal interface {
	speak()
}
type cat string

func (c *cat) speak() {
	fmt.Println("喵喵")
}

type dog string

func (d dog) speak() {
	fmt.Println("汪汪")
}
func main() {
	var a animal = dog("狗")
	a.speak()

	var b cat = cat("猫")
	b.speak()

	(&b).speak()
	var c animal = &b
	c.speak()

	fmt.Println(a)
}
