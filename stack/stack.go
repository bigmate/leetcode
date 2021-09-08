package stack

type Iterable interface {
	Next() (data interface{}, exists bool)
}

type iter struct {
	tmp *node
}

func (i *iter) Next() (data interface{}, exists bool) {
	if i.tmp == nil {
		return nil, false
	}
	current := i.tmp
	i.tmp = current.next
	return current.data, true
}

type Stack interface {
	Push(data interface{})
	Pop() (interface{}, bool)
	Peek() (interface{}, bool)
	Size() int
	Iter() Iterable
}

type stackBuffer struct {
	head *node
	tail *node
	size int
}

func (s *stackBuffer) Iter() Iterable {
	return &iter{tmp: s.head}
}

func New() Stack {
	return &stackBuffer{}
}

func (s *stackBuffer) Size() int {
	return s.size
}

func (s *stackBuffer) Push(data interface{}) {
	n := &node{data: data}
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

func (s *stackBuffer) Pop() (interface{}, bool) {
	if s.tail == nil {
		return nil, false
	}
	if s.tail == s.head {
		s.head = nil
	}
	s.size--
	n := s.tail
	s.tail = n.prev
	return n.data, true
}

func (s *stackBuffer) Peek() (interface{}, bool) {
	if s.tail == nil {
		return nil, false
	}
	return s.tail.data, true
}

type node struct {
	data interface{}
	next *node
	prev *node
}
