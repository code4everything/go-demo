package string

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// LeetCode(id=1071,title=字符串的最大公因子,difficulty=easy)
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	var gcd func(a int, b int) int
	gcd = func(a int, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}

	return str1[:gcd(len(str1), len(str2))]
}

// LeetCode(id=811,title=子域名访问计数,difficulty=easy)
func subdomainVisits(cpdomains []string) []string {
	var helper func(domain string, cnt int)
	domainMap := make(map[string]int)
	helper = func(domain string, cnt int) {
		_, has := domainMap[domain]
		if has {
			domainMap[domain] += cnt
		} else {
			domainMap[domain] = cnt
		}
		idx := strings.Index(domain, ".") + 1
		if idx > 0 {
			helper(domain[idx:], cnt)
		}
	}
	for _, domain := range cpdomains {
		split := strings.Split(domain, " ")
		cnt, _ := strconv.Atoi(split[0])
		helper(split[1], cnt)
	}
	ans, i := make([]string, len(domainMap)), 0
	for k, v := range domainMap {
		ans[i] = strconv.Itoa(v) + " " + k
		i++
	}
	return ans
}

// LeetCode(id=796,title=旋转字符串,difficulty=easy)
func rotateString(A string, B string) bool {
	lenA := len(A)
	lenB := len(B)
	if lenA != lenB {
		return false
	}
	if lenA == 0 {
		return true
	}
	for i := 0; i < lenA; i++ {
		temp, j := i, 0
		for j < lenB {
			if A[temp%lenA] != B[j] {
				break
			}
			temp++
			j++
		}
		if j == lenB {
			return true
		}
	}
	return false
}

// LeetCode(id=748,title=最短完整词,difficulty=easy)
func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlate = strings.ToLower(licensePlate)
	licMap := make([]int, 26)
	for _, v := range licensePlate {
		idx := v - 'a'
		if idx >= 0 && idx < 26 {
			licMap[idx]++
		}
	}
	letterCnt := 0
	for _, v := range licMap {
		letterCnt += v
	}
	word, minSize := "", math.MaxInt32
	for _, v := range words {
		lic, cnt := make([]int, 26), letterCnt
		copy(lic, licMap)
		for _, c := range v {
			idx := c - 'a'
			if lic[idx]--; lic[idx] >= 0 {
				cnt--
			}
		}
		if cnt == 0 {
			if len(v) < minSize {
				minSize = len(v)
				word = v
			}
		}
	}
	return word
}

// LeetCode(id=686,title=重复叠加字符串匹配,difficulty=easy)
func repeatedStringMatch(A string, B string) int {
	s := A
	for i := 1; i <= len(B)/len(A)+2; i++ {
		if strings.Contains(s, B) {
			return i
		}
		s += A
	}
	return -1
}

// LeetCode(id=680,title=验证回文字符串 Ⅱ,difficulty=easy)
func validPalindrome(s string) bool {
	var helper func(s string, left, right, delete int) bool
	helper = func(s string, left, right, delete int) bool {
		for left < right {
			if s[left] == s[right] {
				left++
				right--
			} else if delete > 0 {
				return false
			} else {
				return helper(s, left+1, right, delete+1) || helper(s, left, right-1, delete+1)
			}
		}
		return true
	}
	return helper(s, 0, len(s)-1, 0)
}

// LeetCode(id=665,title=非递减数列,difficulty=easy)
func checkPossibility(nums []int) bool {
	desc, compare, base := 0, nums[0], math.MinInt32
	for i := 1; i < len(nums); i++ {
		if compare > nums[i] {
			desc++
			if nums[i] < base {
				base = compare
			}
		} else {
			if compare > base {
				base = compare
			}
			if base > nums[i] {
				return false
			}
		}
		compare = nums[i]
	}
	return desc < 2
}

// LeetCode(id=657,title=机器人能否返回原点,difficulty=easy)
func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, c := range moves {
		if c == 'U' {
			y++
		} else if c == 'D' {
			y--
		} else if c == 'L' {
			x--
		} else {
			x++
		}
	}
	return x == 0 && y == 0
}

// LeetCode(id=541,title=反转字符串 II,difficulty=easy)
func reverseStr(s string, k int) string {
	var begin, left int
	buf := make([]rune, len(s))
	for i, r := range s {
		if i%(2*k) == 0 {
			begin, left = i, k
			if begin+left-1 >= len(s) {
				left = len(s) - begin
			}
		}
		if left > 0 {
			buf[begin+left-1] = r
			left--
		} else {
			buf[i] = r
		}
	}
	return string(buf)
}

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
