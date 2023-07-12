package heap

type MinHeap struct {
	length int
	data   []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		length: 0,
		data:   make([]int, 0),
	}
}

func (h *MinHeap) Push(val int) {
	h.data = append(h.data, val)
	h.length++
	h.siftUp(h.length - 1)
}

func (h *MinHeap) Pop() int {
	if h.length == 0 {
		return -1
	}
	val := h.data[0]
	h.data[0] = h.data[h.length-1]
	h.data = h.data[:h.length-1]
	h.length--
	h.siftDown(0)
	return val
}

func (h *MinHeap) siftUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[parent] > h.data[index] {
			h.data[parent], h.data[index] = h.data[index], h.data[parent]
			index = parent
		} else {
			break
		}
	}
}

func (h *MinHeap) siftDown(index int) {
	for index < h.length {
		left := index*2 + 1
		right := index*2 + 2
		min := index
		if left < h.length && h.data[left] < h.data[min] {
			min = left
		}
		if right < h.length && h.data[right] < h.data[min] {
			min = right
		}
		if min == index {
			break
		}
		h.data[min], h.data[index] = h.data[index], h.data[min]
		index = min
	}
}
