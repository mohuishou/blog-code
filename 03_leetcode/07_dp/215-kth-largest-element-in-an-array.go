package dp

import "math"

var null = math.MinInt64

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
// 朴素解法，直接维护一个 K 大的数组，替换掉最小的数
func findKthLargest(nums []int, k int) int {
	arr := append([]int{}, nums[:k]...)
	findMin := func(data []int) int {
		minIndex := -1
		min := math.MaxInt64
		for i, n := range data {
			if n < min {
				minIndex = i
				min = n
			}
		}
		return minIndex
	}
	index := findMin(arr)
	for i := k; i < len(nums); i++ {
		if arr[index] < nums[i] {
			arr[index] = nums[i]
			index = findMin(arr)
		}
	}

	return arr[findMin(arr)]
}

// 小顶堆解法
func findKthLargestMinHeap(nums []int, k int) int {
	heap := &heapMin{
		data:   make([]int, k+1),
		length: 0,
	}
	for _, num := range nums {
		if !heap.full() {
			heap.add(num)
		} else if heap.top() < num {
			heap.pop()
			heap.add(num)
		}
	}
	return heap.top()
}

// 堆排序解法
func findKthLargestHeapSort(nums []int, k int) int {
	heap := append([]int{0}, nums...)
	// 1. 建堆
	for i := len(heap) / 2; i > 0; i-- {
		for {
			maxIdx := i
			if heap[i*2] > heap[maxIdx] {
				maxIdx = i * 2
			}
			if i*2+1 < len(heap) && heap[i*2+1] > heap[maxIdx] {
				maxIdx = i*2 + 1
			}
			if maxIdx == i {
				break
			}
			heap[i], heap[maxIdx] = heap[maxIdx], heap[i]
		}
	}

	// 2. 排序
	for count := len(heap) - 1; count > 0; count-- {
		heap[1], heap[count] = heap[count], heap[1]
		i := 1
		for {
			maxIdx := i
			if i*2 < count && heap[i*2] > heap[maxIdx] {
				maxIdx = i * 2
			}
			if i*2+1 < count && heap[i*2+1] > heap[maxIdx] {
				maxIdx = i*2 + 1
			}
			if maxIdx == i {
				break
			}
			heap[i], heap[maxIdx] = heap[maxIdx], heap[i]
			i = maxIdx
		}
	}

	return heap[k]
}
