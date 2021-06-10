package hash

func twoSum(nums []int, target int) []int {
	data := map[int]int{}
	for i, num := range nums {
		if index, ok := data[target-num]; ok {
			return []int{index, i}
		}
		data[num] = i
	}
	return nil
}
