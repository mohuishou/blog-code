package stack

import "math"

// https://leetcode-cn.com/problems/132-pattern/solution/xiang-xin-ke-xue-xi-lie-xiang-jie-wei-he-95gt/
func find132pattern(nums []int) bool {
	k := math.MinInt64
	var stack []int
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < k {
			return true
		}
		j := len(stack) - 1
		for ; j >= 0 && stack[j] < nums[i]; j-- {
			if stack[j] > k {
				k = stack[j]
			}
		}
		stack = stack[:j+1]
		stack = append(stack, nums[i])
	}
	return false
}
