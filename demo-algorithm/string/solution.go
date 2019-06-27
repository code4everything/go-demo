package string

import (
	"sort"
	"strings"
)

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
