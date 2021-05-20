package sort

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	data := mergeIntervals(intervals)
	sort.Sort(data)
	res := data[len(data)-1:]
	for i := len(data) - 1; i >= 0; i-- {
		if res[0][0] <= data[i][1] {
			res[0][0] = min(res[0][0], data[i][0])
			res[0][1] = max(res[0][1], data[i][1])
		} else {
			res = append(mergeIntervals{data[i]}, res...)
		}
	}
	return res
}

type mergeIntervals [][]int

func (m mergeIntervals) Len() int {
	return len(m)
}

func (m mergeIntervals) Less(i, j int) bool {
	return m[i][1] < m[j][1]
}

func (m mergeIntervals) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func min(data ...int) int {
	res := math.MaxInt64
	for _, v := range data {
		if v < res {
			res = v
		}
	}
	return res
}
func max(data ...int) int {
	res := math.MinInt64
	for _, v := range data {
		if v > res {
			res = v
		}
	}
	return res
}
