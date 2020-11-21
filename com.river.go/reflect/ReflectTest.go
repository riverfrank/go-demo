package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u User) Print(prfix string) {
	fmt.Printf("%s:Name is %s,Age is %d", prfix, u.Name, u.Age)
}

func main() {
	var u User
	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		fmt.Println(sf.Tag.Get("json"))
	}
	//
	//u:=User{"张三",20}
	//v:=reflect.ValueOf(u)
	//fmt.Println(v)
	//
	//mPrint:=v.MethodByName("Print")
	//args:=[]reflect.Value{reflect.ValueOf("前缀")}
	//
	//fmt.Println(args)
	//fmt.Println(mPrint.Call(args))
	//
	//
	//
	//t := reflect.TypeOf(u)
	//fmt.Println(t)
	//
	//fmt.Println(t.Kind())
	//
	//for i := 0; i < t.NumField(); i++ {
	//	fmt.Println(t.Field(i).Name)
	//}
}
