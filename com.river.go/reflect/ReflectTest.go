package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{"rivet", 30}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	fmt.Println(t)
	fmt.Println(v)

	fmt.Println(t.Kind())

	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}
}
