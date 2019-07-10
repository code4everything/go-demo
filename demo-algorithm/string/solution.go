package string

import (
	"sort"
	"strings"
)

// LeetCode(id=500,title=键盘行,difficulty=easy)
func findWords(words []string) []string {
	lineMap := make(map[byte]int)
	fillMap(lineMap, []byte{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'}, 1)
	fillMap(lineMap, []byte{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'}, 2)
	fillMap(lineMap, []byte{'z', 'x', 'c', 'v', 'b', 'n', 'm'}, 3)
	ans := make([]string, 0)
	for _, word := range words {
		line, _ := lineMap[word[0]]
		isInLine := true
		for i := 1; i < len(word); i++ {
			if line != lineMap[word[i]] {
				isInLine = false
				break
			}
		}
		if isInLine {
			ans = append(ans, word)
		}
	}
	return ans
}

func fillMap(m map[byte]int, dict []byte, line int) {
	for _, b := range dict {
		m[b] = line
		m[b-32] = line
	}
}

// LeetCode(id=389,title=找不同,difficulty=easy)
func findTheDifference(s string, t string) byte {
	var sum1, sum2 int
	for _, b := range s {
		sum1 += int(b)
	}
	for _, b := range t {
		sum2 += int(b)
	}
	return byte(sum2 - sum1)
}

// LeetCode(id=383,title=赎金信,difficulty=easy)
func canConstruct(ransomNote string, magazine string) bool {
	cnt := make([]int, 26)
	bs := []byte(magazine)
	for _, b := range bs {
		cnt[b-97]++
	}
	bs = []byte(ransomNote)
	for _, b := range bs {
		idx := b - 97
		if cnt[idx] == 0 {
			return false
		}
		cnt[idx]--
	}
	return true
}

// LeetCode(id=290,title=单词规律,difficulty=easy)
func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	size := len(pattern)
	if size != len(strs) {
		return false
	}
	patternMap := make(map[byte]string, size)
	stringMap := make(map[string]byte, size)
	bytes := []byte(pattern)
	for i := 0; i < size; i++ {
		s := strs[i]
		b := bytes[i]
		if p, has := patternMap[b]; has {
			// check p -> s
			if s != p {
				return false
			}
		} else {
			// check s -> p
			if _, has := stringMap[s]; has {
				return false
			}
			stringMap[s] = b
			patternMap[b] = s
		}
	}
	return true
}

// LeetCode(id=58,title=最后一个单词的长度,difficulty=easy)
func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	arr := strings.Split(s, " ")
	if len(arr) == 0 {
		return 0
	}
	return len(arr[len(arr)-1])
}

// LeetCode(id=451,title=根据字符出现频率排序,difficulty=medium)
func frequencySort(s string) string {
	m := make(map[string]int)
	var has bool
	var value int
	// count
	for _, c := range s {
		key := string(c)
		if value, has = m[key]; has {
			m[key] = value + 1
		} else {
			m[key] = 1
		}
	}
	// sort
	p := make(pairs, len(m))
	i := 0
	for k, v := range m {
		p[i] = pair{k, v}
		i++
	}
	sort.Sort(p)
	// merge
	ss := make([]string, len(m))
	i = 0
	for _, a := range p {
		ss[i] = strings.Repeat(a.key, a.value)
		i++
	}
	return strings.Join(ss, "")
}

type pair struct {
	key   string
	value int
}

type pairs []pair

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Less(i, j int) bool {
	return p[i].value > p[j].value
}
