package math

var pos = []int{2, 3, 5}

// LeetCode(id=292,title=Nim 游戏,difficulty=easy)
func canWinNim(n int) bool {
	return n&3 != 0
}

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
