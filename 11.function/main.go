package main

import (
	"11.function/mymath2"
	"fmt"
	"reflect"
)

func main() {
	fmt.Printf("1 + 2 = %v\n", Add(1, 2))
	fmt.Printf("2 - 1 = %v\n", mymath2.Sub(2, 1))
	fmt.Println("=================")

	p1 := new(int)
	p2 := new(string)
	p3 := new([3]int)
	type person struct {
		id   int
		name string
	}
	p4 := new(person)
	fmt.Printf("p1:%v type:%v\n", p1, reflect.TypeOf(p1))
	fmt.Printf("p2:%v type:%v\n", p2, reflect.TypeOf(p2))
	fmt.Printf("p3:%v type:%v\n", p3, reflect.TypeOf(p3))
	fmt.Printf("p4:%v type:%v\n", p4, reflect.TypeOf(p4))

	s1 := make([]int, 3)
	s2 := make(map[string]int, 2)
	fmt.Printf("s1:%v type:%v\n", s1, reflect.TypeOf(s1))
	fmt.Printf("s2:%v type:%v\n", s2, reflect.TypeOf(s2))

}
