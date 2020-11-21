package main

import "fmt"

type Age uint

func (age Age) String(name string) {
	fmt.Println(name, "the age is", age)
}
func main() {
	age := Age(12)
	sum := Age.String
	sum(age, "lucy")

}
