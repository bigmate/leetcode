func exist(board [][]byte, word string) bool {
	s := newScanner(board, word)

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if s.scan(y, x) {
				return true
			}
		}
	}

	return false
}

type scanner struct {
	board [][]byte
	y, x  int
	word  string
	idx   int
}

func (s *scanner) visited(y, x int) bool {
	return s.board[y][x] == '#'
}

func (s *scanner) visit(y, x int) {
	s.idx++
	s.board[y][x] = '#'
}

func (s *scanner) unvisit(y, x int, b byte) {
	s.idx--
	s.board[y][x] = b
}

func (s *scanner) scan(y, x int) bool {
	if s.idx == len(s.word) {
		return true
	}

	if !s.withinBoundary(y, x) || s.visited(y, x) {
		return false
	}

	b := s.board[y][x]

	if s.match(y, x) {
		s.visit(y, x)
		defer s.unvisit(y, x, b)

		return s.scan(y-1, x) || s.scan(y+1, x) || s.scan(y, x+1) || s.scan(y, x-1)
	}

	return false
}

func (s *scanner) match(y, x int) bool {
	return s.board[y][x] == s.word[s.idx]
}

func (s *scanner) withinBoundary(y, x int) bool {
	return y >= 0 && y < s.y && x >= 0 && x < s.x
}

func newScanner(board [][]byte, word string) *scanner {
	return &scanner{
		board: board,
		y:     len(board),
		x:     len(board[0]),
		word:  word,
	}
}
