package main

import "fmt"

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
	memo  map[[2]int]struct{}
	word  string
}

func (s *scanner) visited(y, x int) bool {
	_, ok := s.memo[[...]int{y, x}]
	return ok
}

func (s *scanner) visit(y, x int) {
	s.memo[[...]int{y, x}] = struct{}{}
}

func (s *scanner) unvisit(y, x int) {
	delete(s.memo, [...]int{y, x})
}

func (s *scanner) scan(y, x int) bool {
	if len(s.memo) == len(s.word) {
		return true
	}

	if s.visited(y, x) || !s.withinBoundary(y, x) {
		return false
	}

	if s.match(y, x) {
		s.visit(y, x)
		defer s.unvisit(y, x)

		return s.scan(y-1, x) || s.scan(y+1, x) || s.scan(y, x+1) || s.scan(y, x-1)
	}

	return false
}

func (s *scanner) match(y, x int) bool {
	return s.board[y][x] == s.word[len(s.memo)]
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
		memo:  make(map[[2]int]struct{}, len(word)),
	}
}

func main() {
	board := [][]byte{
		{'C', 'A', 'A'},
		{'A', 'A', 'A'},
		{'B', 'C', 'D'},
	}

	fmt.Println(exist(board, "AAB"))
}
