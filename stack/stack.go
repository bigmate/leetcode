package stack

type Iterable[T any] interface {
	Next() (data T, exists bool)
}

type iter[T any] struct {
	tmp *node[T]
}

func (i *iter[T]) Next() (data T, exists bool) {
	var zeroData T
	if i.tmp == nil {
		return zeroData, false
	}
	current := i.tmp
	i.tmp = current.next
	return current.data, true
}

type Stack[T any] interface {
	Push(data T)
	Pop() (T, bool)
	Peek() (T, bool)
	Size() int
	Iter() Iterable[T]
}

type stackBuffer[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func (s *stackBuffer[T]) Iter() Iterable[T] {
	return &iter[T]{tmp: s.head}
}

func New[T any]() Stack[T] {
	return &stackBuffer[T]{}
}

func (s *stackBuffer[T]) Size() int {
	return s.size
}

func (s *stackBuffer[T]) Push(data T) {
	n := &node[T]{data: data}
	s.size++
	if s.head == nil {
		s.head = n
		s.tail = n
		return
	}
	s.tail.next = n
	n.prev = s.tail
	s.tail = n
}

func (s *stackBuffer[T]) Pop() (T, bool) {
	var zeroData T
	if s.tail == nil {
		return zeroData, false
	}
	if s.tail == s.head {
		s.head = nil
	}
	s.size--
	n := s.tail
	s.tail = n.prev
	return n.data, true
}

func (s *stackBuffer[T]) Peek() (T, bool) {
	var zeroData T
	if s.tail == nil {
		return zeroData, false
	}
	return s.tail.data, true
}

type node[T any] struct {
	data T
	next *node[T]
	prev *node[T]
}
