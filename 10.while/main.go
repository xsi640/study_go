package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("==============")
	i := 0
	for {
		fmt.Println(i)
		i++
		if i > 3 {
			break
		}
	}
	fmt.Println("==============")
	i = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("==============")
	a := []int{1, 2, 3, 4, 5, 6}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)
	fmt.Println("===============")
	for i, v := range a {
		fmt.Printf("%v:%v\n", i, v)
	}
	fmt.Println("===============")
JLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop
			}
			fmt.Printf("i:%v j:%v\n", i, j)
		}
	}
	fmt.Println("========================")
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				goto EXIT
			}
			fmt.Printf("i:%v j:%v\n", i, j)
		}
	}
EXIT:
}
