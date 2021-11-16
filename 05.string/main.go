package main

import "fmt"

func main() {
	v1 := "hello, world!!"
	fmt.Printf("v1: %v, len: %v\n", v1, len(v1))
	v2 := `Search results for "golang"`
	fmt.Printf(v2)
	v3 := `Search results for "Golang":
- Go
- Golang
Golang Programming
`
	fmt.Printf(v3)
	v4 := v1[:5]
	v5 := v1[7:]
	v6 := v1[1:5]
	fmt.Printf(":5: %v\n", v4)
	fmt.Printf("7:: %v\n", v5)
	fmt.Printf("1:5: %v\n", v6)

	fmt.Println("for 遍历字符串")
	for i := 0; i < len(v1); i++ {
		ch := v1[i]
		fmt.Println(i, ch)
	}

	fmt.Println("for range 遍历字符串")
	for i, ch := range v1 {
		fmt.Println(i, string(ch))
	}
}
