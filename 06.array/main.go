package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v1 [8]byte
	var v2 [3][3]int
	var v3 [3][3][3]float64
	var v4 = [3]int{1, 2, 3}
	var v5 = [3]string{"a", "b", "c"}
	var v6 = new([3]string)
	var v7 = [...]int{4, 5, 6}

	fmt.Printf("v1: %v\n", v1)
	fmt.Printf("v2: %v\n", v2)
	fmt.Printf("v3: %v\n", v3)
	fmt.Printf("v4: %v\n", v4)
	fmt.Printf("v5: %v\n", v5)
	fmt.Printf("v6: %v\n", v6)
	fmt.Printf("v7: %v\n", v7)

	fmt.Println("for 遍历数组")
	for i := 0; i < len(v5); i++ {
		fmt.Println(v5[i])
	}
	fmt.Println("for...range 遍历数组")
	for i, v := range v5 {
		fmt.Printf("index:%v value:%v\n", i, v)
	}
	fmt.Println("================")

	months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	q2 := months[3:6]
	fmt.Printf("Q2: %v, type:%v\n", q2, reflect.TypeOf(q2))
	slice1 := make([]int, 5)
	slice2 := make([]int, 5, 10)
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v len:%v cap:%v\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("%v len:%v cap:%v\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("%v len:%v cap:%v\n", slice3, len(slice3), cap(slice3))
	slice4 := append(slice3, 6, 7, 8)
	fmt.Printf("%v len:%v cap:%v\n", slice4, len(slice4), cap(slice4))
	slice5 := append(slice3, []int{6, 7, 8}...)
	fmt.Printf("%v len:%v cap:%v\n", slice5, len(slice5), cap(slice5))
	
}
