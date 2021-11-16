package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 100
	var ptr *int //创建int类型的指针，指针初始化nil
	ptr = &a     //指针指向变量a
	fmt.Printf("打印指针的内存地址：%v\n", ptr)
	fmt.Printf("打印指针指向的值：%v\n", *ptr)

	a1, b1 := 1, 2
	swap(a1, b1)
	fmt.Printf("不使用指针交换数字 %v, %v\n", a1, b1)
	swapPtr(&a1, &b1)
	fmt.Printf("使用指针交换数字 %v, %v \n", a1, b1)
	fmt.Println("=================")
	v1 := 100
	var p *int = &v1
	var fp *float32 = (*float32)(unsafe.Pointer(p))
	*fp = *fp * 10
	fmt.Println(v1)
	fmt.Println("=================")

	arr := [3]int{1, 2, 3}
	ap := &arr //指针ap指向arr
	//ap+偏移量，指向arr的第二个元素
	sp := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ap)) + unsafe.Sizeof(arr[0])))
	*sp += 3 //arr第二个元素+3
	fmt.Println(arr)
	fmt.Println("=================")
}

func swap(a, b int) {
	a, b = b, a
	fmt.Printf("swap %v,%v\n", a, b)
}

func swapPtr(a, b *int) {
	*a, *b = *b, *a
	fmt.Printf("swap %v,%v\n", *a, *b)
}
