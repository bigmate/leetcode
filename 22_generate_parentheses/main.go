package generateParentheses

func generateParenthesis(n int) []string {
	bucket := make([]string, 0, n)
	stack := &stackBuffer{}
	generate(&bucket, stack, n, n)
	return bucket
}

func generate(bucket *[]string, buf ParenthesesStack, left, right int) {
	if left == 0 && right == 0 {
		if buf.WellFormed() {
			*bucket = append(*bucket, buf.String())
		}
	}
	if left > 0 {
		buf.Push('(')
		generate(bucket, buf, left-1, right)
	}
	if right > 0 && buf.LeftCount() > buf.RightCount() {
		buf.Push(')')
		generate(bucket, buf, left, right-1)
	}
	buf.Pop()
}

type ParenthesesStack interface {
	Push(char byte)
	Pop() (byte, bool)
	Peek() (byte, bool)
	String() string // for debugging
	LeftCount() int
	RightCount() int
	Size() int
	WellFormed() bool
}

type stackBuffer struct {
	left  int
	right int
	head  *node
	tail  *node
}

func (s *stackBuffer) Size() int {
	return s.left + s.right
}

func (s *stackBuffer) WellFormed() bool {
	buf := &stackBuffer{}
	n := s.head
	for n != nil {
		switch n.char {
		case '(':
			buf.Push(n.char)
		case ')':
			buf.Pop()
		}
		n = n.next
	}
	return buf.Size() == 0
}

func (s *stackBuffer) LeftCount() int {
	return s.left
}

func (s *stackBuffer) RightCount() int {
	return s.right
}

func (s *stackBuffer) Push(char byte) {
	switch char {
	case '(':
		s.left++
	case ')':
		s.right++
	default:
		panic("invalid byte passed")
	}
	n := &node{char: char}
	if s.head == nil {
		s.head = n
		s.tail = n
		return
	}
	s.tail.next = n
	n.prev = s.tail
	s.tail = n
}

func (s *stackBuffer) Pop() (byte, bool) {
	if s.tail == nil {
		return 0, false
	}
	if s.tail == s.head {
		s.head = nil
	}
	n := s.tail
	s.tail = n.prev
	switch n.char {
	case '(':
		s.left--
	case ')':
		s.right--
	}
	return n.char, true
}

func (s *stackBuffer) Peek() (byte, bool) {
	if s.tail == nil {
		return 0, false
	}
	return s.tail.char, true
}

func (s *stackBuffer) String() string {
	buf := make([]byte, 0, s.Size())
	n := s.head
	for n != nil {
		buf = append(buf, n.char)
		n = n.next
	}
	return string(buf)
}

type node struct {
	char byte
	next *node
	prev *node
}
