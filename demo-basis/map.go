package main

import (
	"fmt"
	"sort"
)

func main() {
	var value int
	var isPresent bool

	// map的初始化
	aMap := make(map[string]int)
	aMap["New Delhi"] = 55
	aMap["Beijing"] = 20
	aMap["Washington"] = 25

	// 获取键值
	value, isPresent = aMap["Beijing"]

	if isPresent {
		fmt.Printf("the value of 'Beijing' in map is: %d\n", value)
	} else {
		fmt.Println("map does not contains 'Beijing'")
	}

	// 删除
	delete(aMap, "Beijing")
	value, isPresent = aMap["Beijing"]

	if isPresent {
		fmt.Printf("the value of 'Beijing' in map is: %d\n", value)
	} else {
		fmt.Println("map does not contains 'Beijing'")
	}

	fmt.Println()
	// map的遍历
	for k, v := range aMap {
		fmt.Printf("%d is value of key '%s'\n", v, k)
	}

	fmt.Println()

	var (
		// 初始化的第二种方式
		barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
			"delta": 87, "echo": 56, "foxtrot": 12,
			"golf": 34, "hotel": 16, "indio": 87,
			"juliet": 65, "kili": 43, "lima": 98}
	)

	fmt.Println("unsorted: ")

	for k, v := range barVal {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}

	fmt.Println()

	// 按key方式排序
	keys := make([]string, len(barVal))
	i := 0
	for k := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("sorted: ")
	for _, k := range keys {
		fmt.Printf("key: %s, value: %d\n", k, barVal[k])
	}

	// 切片
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 1
	}

	fmt.Println(aMap)
	fmt.Println(items)
}
