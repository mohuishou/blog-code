package dp

// 小顶堆
type heapMin struct {
	data   []int
	length int
}

// 向堆中添加一个元素
func (h *heapMin) add(v int) {
	h.length++
	h.data[h.length] = v
	for i := h.length; i/2 > 0 && h.data[i/2] > h.data[i]; i = i / 2 {
		h.data[i/2], h.data[i] = h.data[i], h.data[i/2]
	}
}

func (h *heapMin) top() int {
	return h.data[1]
}

func (h *heapMin) full() bool {
	return h.length == len(h.data)-1
}

// 弹出堆顶元素
func (h *heapMin) pop() int {
	if h.length == 0 {
		return null
	}

	v := h.data[1]
	h.data[1] = h.data[h.length]

	i := 1
	for {
		minIdx := i
		if i*2 < h.length && h.data[i] > h.data[2*i] {
			minIdx = 2 * i
		}
		if i*2+1 < h.length && h.data[minIdx] > h.data[2*i+1] {
			minIdx = 2*i + 1
		}
		// 说明当前已经是最小的值了
		if minIdx == i {
			break
		}
		h.data[i], h.data[minIdx] = h.data[minIdx], h.data[i]
		i = minIdx
	}
	h.length--
	return v
}
