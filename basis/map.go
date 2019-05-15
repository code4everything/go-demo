package main

import (
	"fmt"
	"sort"
)

func main() {
	var value int
	var isPresent bool

	aMap := make(map[string]int)
	aMap["New Delhi"] = 55
	aMap["Beijing"] = 20
	aMap["Washington"] = 25

	value, isPresent = aMap["Beijing"]

	if isPresent {
		fmt.Printf("the value of 'Beijing' in map is: %d\n", value)
	} else {
		fmt.Println("map does not contains 'Beijing'")
	}

	delete(aMap, "Beijing")
	value, isPresent = aMap["Beijing"]

	if isPresent {
		fmt.Printf("the value of 'Beijing' in map is: %d\n", value)
	} else {
		fmt.Println("map does not contains 'Beijing'")
	}

	fmt.Println()
	for k, v := range aMap {
		fmt.Printf("%d is value of key '%s'\n", v, k)
	}

	fmt.Println()

	var (
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
}
