package math

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
