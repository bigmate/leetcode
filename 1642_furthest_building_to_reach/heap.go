package main

type HeapInt struct {
	s    []int
	less func(i, j int) bool
}

func (h *HeapInt) Len() int {
	return len(h.s)
}

func (h *HeapInt) Less(i, j int) bool {
	return h.less(i, j)
}

func (h *HeapInt) Swap(i, j int) {
	h.s[i], h.s[j] = h.s[j], h.s[i]
}

func (h *HeapInt) Push(x any) {
	h.s = append(h.s, x.(int))
}

func (h *HeapInt) Pop() any {
	el := h.s[h.Len()-1]
	h.s = h.s[:h.Len()-1]
	return el
}

func NewMaxHeap(size int) *HeapInt {
	h := &HeapInt{
		s: make([]int, 0, size),
	}

	var less func(i, j int) bool
	less = func(i, j int) bool {
		return h.s[i] > h.s[j]
	}

	h.less = less

	return h
}
