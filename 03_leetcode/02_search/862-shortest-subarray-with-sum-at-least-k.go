package search

import "math"

type shortestSubarrayStack struct {
	data    [][2]int
	current int
}

func (s *shortestSubarrayStack) push(sum, index int) {
	s.data[s.current] = [2]int{sum, index}
	s.current++
}

func (s *shortestSubarrayStack) pop() (sum, index int) {
	sum, index = s.last()
	if sum != math.MinInt32 {
		s.current--
	}
	return
}

func (s *shortestSubarrayStack) last() (sum, index int) {
	if s.current == 0 {
		return math.MinInt32, -1
	}
	return s.data[s.current-1][0], s.data[s.current-1][1]
}

// 查找小于 target 的最小值
func (s *shortestSubarrayStack) find(target int) (index int) {
	left, right := 0, s.current-1
	for left <= right {
		mid := left + (right-left)/2
		if s.data[mid][0] < target {
			if s.data[mid+1][0] > target {
				return s.data[mid][1]
			}
			left = mid + 1
		}
		if s.data[mid][0] == target {
			return s.data[mid][1]
		}
		if s.data[mid][0] > target {
			right = mid - 1
		}
	}
	return -2
}

func shortestSubarray(nums []int, k int) int {
	// 用于保存前缀和，以及当前前缀和的索引地址
	// 保存单调递增的前缀和
	// 例如现在有 (1,2) (3,3) 两个数据，下一个前缀和是 (2,4)
	// 那么 sums 就会变成 (1,2) (2,4)
	sums := shortestSubarrayStack{data: make([][2]int, len(nums)+1)}
	sums.push(0, -1)
	min := math.MaxInt32

	var sum int
	for i, v := range nums {
		sum += v
		s, _ := sums.last()
		for s > sum {
			sums.pop()
			s, _ = sums.last()
		}
		sums.push(sum, i)
		if index := sums.find(sum - k); index != -2 {
			if i-index < min {
				min = i - index
			}
		}
	}
	if min == math.MaxInt32 {
		return -1
	}
	return min
}
