package search

func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		start, end := 0, len(row)-1
		for start <= end {
			mid := start + (end-start)/2
			if row[mid] < target {
				start = mid + 1
			} else if row[mid] == target {
				return true
			} else {
				end = mid - 1
			}
		}
	}
	return false
}
