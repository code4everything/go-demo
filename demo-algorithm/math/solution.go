package math

import "strconv"

// LeetCode(id=400,title=第N个数字,difficulty=easy)
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	var lens int
	left := 1
	right := n
	for left < right {
		mid := left + (right-left)>>1
		if mid < 10 {
			lens = mid
		} else if mid < 100 {
			lens = 9 + 2*(mid-9)
		} else {
			lens = 189
			div1 := 99
			div2 := 999
			i := 3
			for mid > div1 {
				if mid <= div2 {
					lens += (mid - div1) * i
					break
				} else {
					lens += (div2 - div1) * i
				}
				div1 = div2
				div2 = div2*10 + 9
				i++
			}
		}
		if lens == n {
			return mid
		}
		if lens < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	b := []byte(strconv.Itoa(left))
	return int(b[n-lens] - '0')
}

// LeetCode(id=342,title=4的幂,difficulty=easy)
func isPowerOfFour(num int) bool {
	// 二进制 1010 1010 1010 1010 1010 1010 1010 1010 的十进制表示
	even := 2863311530
	if num&even != 0 {
		return false
	}
	var cnt = 0
	for num > 0 {
		if num&1 == 1 {
			if cnt == 1 {
				return false
			}
			cnt++
		}
		num >>= 2
	}
	return cnt == 1
}

// LeetCode(id=292,title=Nim 游戏,difficulty=easy)
func canWinNim(n int) bool {
	return n&3 != 0
}

var pos = []int{2, 3, 5}

// LeetCode(id=263,title=丑数,difficulty=easy)
func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	for i := 0; i < 3 && num > 1; {
		if num%pos[i] == 0 {
			num /= pos[i]
		} else {
			i++
		}
	}
	return num == 1
}

// LeetCode(id=258,title=各位相加,difficulty=easy)
func addDigits(num int) int {
	if num < 10 {
		return num
	}
	mod := num % 9
	if mod == 0 {
		return 9
	}
	return mod
}

// LeetCode(id=231,title=2的幂,difficulty=easy)
func isPowerOfTwo(n int) bool {
	var cnt = 0
	for n > 0 {
		if n&1 == 1 {
			if cnt == 1 {
				return false
			}
			cnt++
		}
		n >>= 1
	}
	return cnt == 1
}

// LeetCode(id=171,title=Excel表列序号,difficulty=easy)
func convertToTitle(n int) string {
	bytes := make([]byte, 0)
	for n > 0 {
		n--
		bytes = append(bytes, byte(65+n%26))
		n /= 26
	}
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}
