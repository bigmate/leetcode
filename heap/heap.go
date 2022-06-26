package heap

import (
	"errors"
	"golang.org/x/exp/constraints"
)

var EmptyHeapErr = errors.New("heap: empty heap")

type Heap[T any] interface {
	Push(T)
	Pop() (T, error)
	Peek() (T, error)
	Len() int
}

type hp[T constraints.Ordered] struct {
	store []T
	less  func(i, j int) bool
}

func NewMaxHeap[T constraints.Ordered](size int) Heap[T] {
	h := &hp[T]{
		store: make([]T, 0, size),
	}

	var less func(i, j int) bool
	less = func(i, j int) bool {
		return h.store[i] > h.store[j]
	}

	h.less = less

	return h
}

func NewMinHeap[T constraints.Ordered](size int) Heap[T] {
	h := &hp[T]{
		store: make([]T, 0, size),
	}

	var less func(i, j int) bool
	less = func(i, j int) bool {
		return h.store[i] < h.store[j]
	}

	h.less = less

	return h
}

func (h *hp[T]) Push(el T) {
	h.store = append(h.store, el)
	h.up(h.Len() - 1)
}

func (h *hp[T]) Pop() (T, error) {
	if h.Len() == 0 {
		var zeroVal T
		return zeroVal, EmptyHeapErr
	}

	n := h.Len() - 1

	h.swap(0, n)
	h.down(0, n)

	el := h.store[n]
	h.store = h.store[:n]

	return el, nil
}

func (h *hp[T]) Peek() (T, error) {
	if h.Len() == 0 {
		var zeroVal T
		return zeroVal, EmptyHeapErr
	}

	return h.store[0], nil
}

func (h *hp[T]) Len() int {
	return len(h.store)
}

func (h *hp[T]) swap(i, j int) {
	h.store[i], h.store[j] = h.store[j], h.store[i]
}

func (h *hp[T]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.less(j, i) {
			break
		}
		h.swap(i, j)
		j = i
	}
}

func (h *hp[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.less(j, i) {
			break
		}
		h.swap(i, j)
		i = j
	}
	return i > i0
}
