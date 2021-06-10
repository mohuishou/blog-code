package dp

func findTargetSumWays(nums []int, target int) int {
	data := map[int]int{0: 1}
	for _, num := range nums {
		nData := map[int]int{}
		for sum, count := range data {
			nData[sum+num] += count
			nData[sum-num] += count
		}
		data = nData
	}
	return data[target]
}
