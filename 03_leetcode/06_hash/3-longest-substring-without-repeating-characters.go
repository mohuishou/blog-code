package hash

func lengthOfLongestSubstring(s string) int {
	var (
		start, end, max int
		data            = map[byte]int{}
	)
	for ; end < len(s); end++ {
		b := s[end]
		i, ok := data[b]
		if ok {
			for j := i; j >= start && j <= i; j-- {
				delete(data, s[j])
			}
			start = i + 1
		} else if max < end-start+1 {
			max = end - start + 1
		}
		data[b] = end
	}
	return max
}
