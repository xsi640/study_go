package main

import "fmt"

func main() {
	if true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	fmt.Println("===========")
	score := 100
	switch {
	case score >= 90:
		fmt.Println("score A")
	case score >= 80 && score < 90:
		fmt.Println("score B")
	case score >= 70 && score < 80:
		fmt.Println("score C")
	case score >= 60 && score < 70:
		fmt.Println("score D")
	default:
		fmt.Println("score F")
	}

	fmt.Println("==============")
	score = 60
	switch score {
	case 90, 100:
		fmt.Println("score A")
	case 80:
		fmt.Println("score B")
	case 70:
		fmt.Println("score C")
	case 60:
		fallthrough //强制执行下面的case语句
	case 65:
		fmt.Println("score D")
	default:
		fmt.Println("score F")
	}
}
