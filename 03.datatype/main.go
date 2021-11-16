package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v1 int
	v1 = 1
	fmt.Printf("v1 = %d\n", v1)
	var (
		v2 int
		v3 string
	)
	v2 = 2
	v3 = "hello"
	fmt.Printf("v2 = %d\n", v2)
	fmt.Printf("v3 = %s\n", v3)
	fmt.Print("====\n\n")

	//没有指定的值时，会自动赋默认值
	var d1 int
	var d2 string
	var d3 bool
	var d4 [10]int
	var d5 struct {
		f float64
	}
	var d6 *int
	var d7 map[string]int
	var d8 func(a int) int
	fmt.Printf("d1 = %v\n", d1)
	fmt.Printf("d2 = %v\n", d2)
	fmt.Printf("d3 = %v\n", d3)
	fmt.Printf("d4 = %v\n", d4)
	fmt.Printf("d5 = %v\n", d5)
	fmt.Printf("d6 = %v\n", d6)
	fmt.Printf("d7 = %v\n", d7)
	fmt.Printf("d8 = %v\n", d8)
	fmt.Print("====\n\n")

	var f1 int = 10
	var f2 = "Hello"
	f3 := 10   //类型推断，自动推断为int
	f4 := 10.0 //自动推断为float64
	fmt.Printf("f1 = %v type:%s\n", f1, reflect.TypeOf(f1))
	fmt.Printf("f2 = %v type:%s\n", f2, reflect.TypeOf(f2))
	fmt.Printf("f3 = %v type:%s\n", f3, reflect.TypeOf(f3))
	fmt.Printf("f4 = %v type:%s\n", f4, reflect.TypeOf(f4))
	fmt.Print("====\n\n")

	//多重赋值
	var g1, g2 = 1, 2
	fmt.Printf("g1 = %v\n", g1)
	fmt.Printf("g2 = %v\n", g2)
	g1, g2 = g2, g1 //交换g1和g2的值
	fmt.Printf("g1 = %v\n", g1)
	fmt.Printf("g2 = %v\n", g2)
	_, nickname := GetName() //利用多重赋值将函数GetName返回的两个类型，第一个返回值忽略、第二个返回值给nickname
	fmt.Printf("nickname=%v\n", nickname)
	fmt.Print("====\n\n")

	const Pi = 3.14159265358979323846
	const zero = 0.0
	const (
		size int64 = 1024
		eof        = -1
	)
	const u, v float32 = 0, 3
	const a, b, c = 3, 4, "foo"
	//u = 2 常量不允许赋值
	fmt.Printf("Pi = %v\n", Pi)
	fmt.Printf("zero = %v\n", zero)
	fmt.Printf("size = %v\n", size)
	fmt.Printf("eof = %v\n", eof)
	fmt.Printf("u, v = %v, %v\n", u, v)
	fmt.Printf("a, b, c = %v, %v, %v\n", a, b, c)
	fmt.Print("====\n\n")

	//预定义常量, ture, false, iota
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	//ioa第二次出现const, iota初始化0
	const (
		z1 = iota
		z2 = iota
		z3 = iota
	)
	fmt.Printf("c1, c2, c3 = %v, %v, %v\n", c1, c2, c3)
	fmt.Printf("z1, z2, z3 = %v, %v, %v\n", z1, z2, z3)
	const (
		x1 = iota * 2
		x2 //后面一样，可省略
		x3
	)
	fmt.Printf("x1, x2, x3 = %v, %v, %v\n", x1, x2, x3)
	fmt.Print("====\n\n")

	//go没有enum定义，通常使用const来定义枚举。
	type Sex uint
	const (
		Male Sex = iota
		Female
	)
	fmt.Printf("Male = %v, Female = %v\n", Male, Female)
	fmt.Print("====\n\n")
}

func GetName() (username, nickname string) {
	return "su", "yang"
}
