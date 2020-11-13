package main

import "fmt"

type Handler interface {
	Do(k string, v int)
}

func Each(m map[string]int, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

type welcome string

func (w *welcome) Do(name string, age int) {
	fmt.Println(w, "welcome ", name, " your age is ", age)
}
func (w welcome) selfInfo(k string, v int) {
	fmt.Printf("%s,我叫%s,今年%d岁\n", w, k, v)
}

type HandlerFunc func(k string, v int)

func (f HandlerFunc) Do(k string, v int) {
	f(k, v)
}

func EachFunc(m map[string]int, f func(k string, v int)) {
	Each(m, HandlerFunc(f))
}

func selfInfo(k string, v int) {
	fmt.Printf("大家好,我叫%s,今年%d岁\n", k, v)
}

func main() {
	nameAge := make(map[string]int)
	nameAge["river"] = 34
	nameAge["lucy"] = 55
	var aa = welcome("bb")
	Each(nameAge, &aa)
	handlerFunc := HandlerFunc(aa.selfInfo)
	Each(nameAge, handlerFunc)

	EachFunc(nameAge, selfInfo)
}
