package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	v1 := uint(16) //自动将int转换成uint
	v2 := int8(v1) //转换成int8
	v3 := uint16(v2)
	v4a := 99.99
	v4 := int(v4a)   //float64转int
	v5 := string(65) //int转字符串
	v6 := string(30028)
	v7 := string([]byte{'h', 'e', 'l', 'l', 'o'})
	v8 := string([]rune{0x5b66, 0x9662, 0x541b}) //rune是int32的别名
	fmt.Printf("v1 = %v, type:%v\n", v1, reflect.TypeOf(v1))
	fmt.Printf("v2 = %v, type:%v\n", v2, reflect.TypeOf(v2))
	fmt.Printf("v3 = %v, type:%v\n", v3, reflect.TypeOf(v3))
	fmt.Printf("v4 = %v, type:%v to type:%v\n", v4, reflect.TypeOf(v4a), reflect.TypeOf(v4))
	fmt.Printf("v5 = %v, type:%v\n", v5, reflect.TypeOf(v5))
	fmt.Printf("v6 = %v, type:%v\n", v6, reflect.TypeOf(v6))
	fmt.Printf("v7 = %v, type:%v\n", v7, reflect.TypeOf(v7))
	fmt.Printf("v8 = %v, type:%v\n", v8, reflect.TypeOf(v8))
	fmt.Println("================")

	d1, _ := strconv.Atoi("100") //字符串转int
	fmt.Printf("d1 = %v, type:%v\n", d1, reflect.TypeOf(d1))
	d2 := strconv.Itoa(100) //int转字符串
	fmt.Printf("d1 = %v, type:%v\n", d2, reflect.TypeOf(d2))
	d3, _ := strconv.ParseBool("false")
	fmt.Printf("d3 = %v, type:%v\n", d3, reflect.TypeOf(d3))
	d4, _ := strconv.ParseInt("100", 10, 64) //转10进制数字，64位
	fmt.Printf("d4 = %v, type:%v\n", d4, reflect.TypeOf(d4))
	d5, _ := strconv.ParseFloat("99.99", 64)
	fmt.Printf("d5 = %v, type:%v\n", d5, reflect.TypeOf(d5))
	d6a := strconv.Quote("Hello, 世界!")//添加引号
	d6 := strconv.QuoteToASCII(d6a)//转asic编码
	fmt.Printf("d6 = %v, type:%v\n", d6, reflect.TypeOf(d6))
}
