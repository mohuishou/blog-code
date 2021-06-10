package hash

import "math"

func minWindow(s string, t string) string {
	data := map[string][]int{}
	countMap := map[string]int{}
	for _, s := range t {
		data[string(s)] = []int{}
		countMap[string(s)]++
	}

	for i, v := range s {
		if _, ok := data[string(v)]; ok {
			data[string(v)] = append(data[string(v)], i)
		}
	}
	start, end := 0, -1
	minSubString := math.MaxInt64

Loop:
	for {
		min, max := math.MaxInt64, math.MinInt64
		minIndex := ""
		for i, arr := range data {
			if len(arr) < countMap[i] {
				break Loop
			}
			if arr[0] < min {
				min = arr[0]
				minIndex = i
			}
			if arr[countMap[i]-1] > max {
				max = arr[countMap[i]-1]
			}
		}
		data[minIndex] = data[minIndex][1:]
		if max-min < minSubString {
			end = max
			start = min
			minSubString = max - min
		}
	}
	return s[start : end+1]
}
