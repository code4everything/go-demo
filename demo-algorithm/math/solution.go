package math

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// LeetCode(id=754,title=到达终点数字,difficulty=easy)
func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	sum, res := 0, 0
	for target > sum || (sum-target)%2 != 0 {
		res++
		sum += res
	}
	return res
}

// LeetCode(id=717,title=1比特与2比特字符,difficulty=easy)
func isOneBitCharacter(bits []int) bool {
	i, size := 0, len(bits)
	for ; i < size-1; i++ {
		if bits[i] == 1 {
			i++
		}
	}
	return i < size
}

// LeetCode(id=697,title=数组的度,difficulty=easy)
func findShortestSubArray(nums []int) int {
	starts, ends, times := make([]int, 50000), make([]int, 50000), make([]int, 50000)
	for i, v := range nums {
		if starts[v] == 0 {
			starts[v] = i + 1
		}
		ends[v] = i + 1
		times[v]++
	}
	max, minLen := 1, 49999
	for i, v := range times {
		if v >= max {
			tmp := ends[i] - starts[i]
			if v > max {
				max = v
				minLen = tmp
			} else if tmp < minLen {
				minLen = tmp
			}
		}
	}
	return minLen + 1
}

// LeetCode(id=693,title=交替位二进制数,difficulty=easy)
func hasAlternatingBits(n int) bool {
	k := int64(n ^ (n >> 1))
	return k&(k+1) == 0
}

// LeetCode(id=674,title=最长连续递增序列,difficulty=easy)
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max, tmp := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			tmp++
		} else {
			if tmp > max {
				max = tmp
			}
			tmp = 1
		}
	}
	if tmp > max {
		return tmp
	}
	return max
}

// LeetCode(id=645,title=错误的集合,difficulty=easy)
func findErrorNums(nums []int) []int {
	size := len(nums) + 1
	bucket := make([]int, size)
	for _, v := range nums {
		bucket[v]++
	}
	ans := make([]int, 2)
	for i := 1; i < size; i++ {
		if bucket[i] == 0 {
			ans[1] = i
		} else if bucket[i] == 2 {
			ans[0] = i
		}
	}
	return ans
}

// LeetCode(id=643,title=子数组最大平均数 I,difficulty=easy)
func findMaxAverage(nums []int, k int) float64 {
	i, sum, size := 0, 0, len(nums)
	for i < k {
		sum += nums[i]
		i++
	}
	max := sum
	for i < size {
		sum = sum - nums[i-k] + nums[i]
		if sum > max {
			max = sum
		}
		i++
	}
	return float64(max) / float64(k)
}

// LeetCode(id=633,title=平方数之和,difficulty=easy)
func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		val := left*left + right*right
		if val == c {
			return true
		} else if val > c {
			right--
		} else {
			left--
		}
	}
	return false
}

// LeetCode(id=628,title=三个数的最大乘积,difficulty=easy)
func maximumProduct(nums []int) int {
	max1, max2, max3 := -1001, -1001, -1001
	min1, min2 := 1001, 1001
	for _, v := range nums {
		if v > max1 {
			max3, max2, max1 = max2, max1, v
		} else if v > max2 {
			max3, max2 = max2, v
		} else if v > max3 {
			max3 = v
		}
		if v < min1 {
			min2, min1 = min1, v
		} else if v < min2 {
			min2 = v
		}
	}
	first, last := min1*min2, max2*max3
	if first > last {
		return first * max1
	}
	return last * max1
}

// LeetCode(id=605,title=种花问题,difficulty=easy)
func canPlaceFlowers(flowerbed []int, n int) bool {
	i, size, feed := 0, len(flowerbed), false
	high := size - 1
	for i < size && n > 0 {
		feed = flowerbed[i] == 0
		if feed {
			if i == 0 && i < high {
				feed = flowerbed[i+1] == 0
			} else if i < high {
				feed = flowerbed[i-1] == 0 && flowerbed[i+1] == 0
			} else if i != 0 {
				feed = flowerbed[i-1] == 0
			}
		}
		if feed {
			flowerbed[i] = 1
			n--
		}
		i++
	}
	return n == 0
}

// LeetCode(id=589,title=范围求和 II,difficulty=easy)
func maxCount(m int, n int, ops [][]int) int {
	minX := m
	minY := n
	for _, xy := range ops {
		if xy[0] < minX {
			minX = xy[0]
		}
		if xy[1] < minY {
			minY = xy[1]
		}
	}
	return minY * minX
}

// LeetCode(id=594,title=最长和谐子序列,difficulty=easy)
func findLHS(nums []int) int {
	lhs := make(map[int][]int)
	unique := make(map[int]bool)
	for _, v := range nums {
		if lh, has := lhs[v]; has {
			lh[0]++
			lh[1]++
		} else {
			lh = make([]int, 2)
			lh[0], lh[1] = 1, 1
			lhs[v] = lh
			unique[v] = true
		}
		less := v - 1
		if lh, has := lhs[less]; has {
			lh[1]++
			unique[less] = false
		}
		great := v + 1
		if lh, has := lhs[great]; has {
			lh[0]++
			unique[great] = false
		}
	}
	max := 0
	for k, v := range lhs {
		if unique[k] {
			continue
		}
		if v[0] > max {
			max = v[0]
		}
		if v[1] > max {
			max = v[1]
		}
	}
	return max
}

// LeetCode(id=581,title=最短无序连续子数组,difficulty=easy)
func findUnsortedSubarray(nums []int) int {
	copier := make([]int, len(nums))
	copy(copier, nums)
	sort.Ints(copier)
	i, j := 0, len(nums)-1
	head, last := false, false
	for i < j {
		if head && last {
			break
		}
		if !head {
			if nums[i] == copier[i] {
				i++
			} else {
				head = true
			}
		}
		if !last {
			if nums[j] == copier[j] {
				j--
			} else {
				last = true
			}
		}
	}
	if i >= j {
		return 0
	}
	return j - i + 1
}

// LeetCode(id=575,title=分糖果,difficulty=easy)
func distributeCandies(candies []int) int {
	types := make(map[int]int)
	for _, v := range candies {
		types[v] = 1
	}
	max := len(candies) >> 1
	size := len(types)
	if size > max {
		return max
	}
	return size
}

// LeetCode(id=532,title=数组中的K-diff数对,difficulty=easy)
func findPairs(nums []int, k int) int {
	used := make(map[int]int)
	sort.Ints(nums)
	ans, i, j := 0, 0, 1
	for j < len(nums) {
		diff := nums[j] - nums[i]
		if diff < k {
			j++
		} else if diff > k {
			i++
			if i == j {
				j++
			}
		} else {
			if _, has := used[nums[i]]; !has {
				ans++
				used[nums[i]] = 1
			}
			i++
			j++
		}
	}
	return ans
}

// LeetCode(id=504,title=七进制数,difficulty=easy)
func convertToBase7(num int) string {
	return strconv.FormatInt(int64(num), 7)
}

// LeetCode(id=496,title=下一个更大元素 I,difficulty=easy)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	size := len(nums2) - 1
	for i, v := range nums1 {
		ng := -1
		p := size
		for nums2[p] != v {
			if nums2[p] > v {
				ng = nums2[p]
			}
			p--
		}
		nums1[i] = ng
	}
	return nums1
}

// LeetCode(id=492,title=构造矩形,difficulty=easy)
func constructRectangle(area int) []int {
	high := int(math.Sqrt(float64(area)))
	for area%high != 0 {
		high--
	}
	return []int{area / high, high}
}

// LeetCode(id=476,title=数字的补数,difficulty=easy)
func findComplement(num int) int {
	base := 1
	for base < num {
		base = (base << 1) + 1
	}
	return num ^ base
}

// LeetCode(id=475,title=供暖器,difficulty=easy)
func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	low, high := 0, 1000000000
	for low < high {
		mid := low + (high-low)/2
		hIdx := 0
		left, right := heaters[hIdx]-mid, heaters[hIdx]+mid
		var i int
		for i = 0; i < len(houses); {
			if houses[i] < left {
				break
			} else if houses[i] > right {
				hIdx++
				if hIdx >= len(heaters) {
					break
				}
				left, right = heaters[hIdx]-mid, heaters[hIdx]+mid
				continue
			}
			i++
		}
		if i == len(houses) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

var watches []string

// LeetCode(id=401,title=进制手表,difficulty=easy)
func readBinaryWatch(num int) []string {
	if num == 0 {
		return []string{"0:00"}
	}
	watches = make([]string, 0)
	watchHelper(num, make([]int, 10), 0)
	return watches
}

func watchHelper(rec int, bits []int, start int) {
	end := 10 - rec
	for start <= end {
		bits[start] = 1
		if rec > 1 {
			// 回溯遍历
			watchHelper(rec-1, bits, start+1)
		} else {
			// 计算时间
			hour := 0
			var i uint32 = 0
			for ; i < 4; i++ {
				hour += bits[i] << (3 - i)
			}
			minute := 0
			for ; i < 10; i++ {
				minute += bits[i] << (9 - i)
			}
			if hour < 12 && minute < 60 {
				watches = append(watches, fmt.Sprintf("%d:%0.2d", hour, minute))
			}
		}
		// 选值结束，进行下一轮选值
		bits[start] = 0
		start++
	}
}

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
			return mid % 10
		}
		if lens < n {
			if left+1 == right {
				left = right
				break
			}
			left = mid
		} else {
			right = mid
		}
	}
	b := []byte(strconv.Itoa(left))
	return int(b[n-lens-1] - '0')
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
