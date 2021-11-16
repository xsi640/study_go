package main

import (
	"fmt"
	"sort"
)

func main() {
	var testMap map[string]int
	testMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	testMap["five"] = 5
	testMap["a"] = 6
	testMap["z"] = 9
	fmt.Printf("testMap: %v\n", testMap)
	k := "two"
	v, ok := testMap[k]
	if ok {
		fmt.Printf("key:%v:%v\n", k, v)
	} else {
		fmt.Printf("not found.\n")
	}

	fmt.Println("for遍历map")
	for k, v := range testMap {
		fmt.Printf("%v:%v\n", k, v)
	}
	fmt.Println("键值对掉")
	invMap := make(map[int]string, len(testMap))
	for k, v := range testMap {
		invMap[v] = k
	}
	fmt.Printf("invMap:%v\n", invMap)
	fmt.Println("按key排序")
	keys := make([]string, 0)
	for k, _ := range testMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	sortedMap := make(map[string]int, len(testMap))
	for _, k := range keys {
		sortedMap[k] = testMap[k]
	}
	fmt.Println("sortedMap:")
	for k, v = range sortedMap {
		fmt.Printf("%v:%v\n", k, v)
	}
}
