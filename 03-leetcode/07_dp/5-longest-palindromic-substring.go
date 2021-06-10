package dp

func longestPalindrome(s string) string {
	var sub string
	for i := range s {
		// 先查找是否存在相同字符串，如果存在就作为一个整体
		k := i + 1
		for ; k < len(s) && s[k] == s[i]; k++ {
		}
		k = k - 1

		// 从中心向外扩散，查找是否相等
		j := 1
		for ; k+j < len(s) && i-j >= 0 && s[k+j] == s[i-j]; j++ {
		}
		j = j - 1
		if len(sub) < (k+j)-(i-j)+1 {
			sub = s[i-j : k+j+1]
		}
	}
	return sub
}
