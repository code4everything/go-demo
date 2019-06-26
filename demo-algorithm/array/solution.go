package array

// LeetCode(id=35,title=搜索插入位置,difficulty=easy)
func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		idx := i + ((j - i) >> 1)
		val := nums[idx]
		if val == target {
			return idx
		} else if val > target {
			j = idx
		} else {
			i = idx
		}
	}
	return i
}

// LeetCode(id=1,title=两数字和,difficulty=medium)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for idx, val := range nums {
		diff := target - val
		if i, has := m[diff]; has {
			return []int{i, idx}
		}
		m[val] = idx
	}
	return nil
}
