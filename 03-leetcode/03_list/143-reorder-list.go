package list

func reorderList(head *ListNode) {
	var data []*ListNode
	for head != nil {
		data = append(data, head)
		next := head.Next
		head.Next = nil
		head = next
	}

	start, end := 0, len(data)-1
	for start < end {
		data[start].Next = data[end]
		if start+1 < end {
			data[end].Next = data[start+1]
		}
		start++
		end--
	}
}
