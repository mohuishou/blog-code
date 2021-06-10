package search

import "math"

// 迭代法
func findPeakElement(nums []int) int {
	nums = append(nums, -math.MaxInt32)
	nums = append([]int{-math.MaxInt32}, nums...)
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] && nums[i] > nums[i-1] {
			return i - 1
		}
	}
	return 0
}

// findPeakElement2 2分法
func findPeakElement2(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if mid+1 >= len(nums) || nums[mid] > nums[mid+1] {
			if mid == 0 || nums[mid] > nums[mid-1] {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return 0
}
