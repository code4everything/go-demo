package string

import (
	"fmt"
	"testing"
)

func TestISubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"asak", "ahbgdclkjlaflkshflakjfoieofalksdlkflsks"}, true},
	}
	fmt.Println(tests)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanConstruct(t *testing.T) {
	canConstruct("a", "ab")
}

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
