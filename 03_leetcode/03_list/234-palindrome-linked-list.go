package list

func isPalindrome(head *ListNode) bool {
	n := head
	var stack []int
	for n != nil {
		stack = append(stack, n.Val)
		n = n.Next
	}
	for i := 0; i < len(stack)-1-i; i++ {
		if stack[i] != stack[len(stack)-1-i] {
			return false
		}
	}
	return true
}
