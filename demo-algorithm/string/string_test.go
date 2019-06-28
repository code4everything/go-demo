package string

import (
	"fmt"
	"testing"
)

func TestWordPattern(t *testing.T) {
	fmt.Println("true\t", wordPattern("abba", "dog cat cat dog"))
	fmt.Println("false\t", wordPattern("abba", "dog cat cat fish"))
	fmt.Println("false\t", wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println("false\t", wordPattern("abba", "dog dog dog dog"))
}

func TestFrequencySort(t *testing.T) {
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("aAbb"))
}
