package string

import (
	"fmt"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("aAbb"))
}
