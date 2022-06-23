package main

import (
	"fmt"
	"math/rand"
)

func longestIncreasingPath(mx [][]int) int {
	m := make(Memo, len(mx)+len(mx[0]))

	for y := 0; y < len(mx); y++ {
		for x := 0; x < len(mx[y]); x++ {
			dfs(m, mx, x, y, 1)
		}
	}

	return m.Max()
}

type Memo map[[2]int]int

func (m Memo) Store(x, y, nl int) {
	ol := m[[...]int{x, y}]
	if ol < nl {
		m[[...]int{x, y}] = nl
	}
}

func (m Memo) Get(x, y int) int {
	return m[[...]int{x, y}]
}

func (m Memo) Max() int {
	var v int

	for _, c := range m {
		if v < c {
			v = c
		}
	}

	return v
}

func dfs(memo Memo, mx [][]int, x, y, l int) {
	if memo.Get(x, y) > l {
		return
	}

	w, h := len(mx[0]), len(mx)
	c := mx[y][x]

	if w > x+1 && mx[y][x+1] > c {
		dfs(memo, mx, x+1, y, l+1)
	}

	if x-1 >= 0 && mx[y][x-1] > c {
		dfs(memo, mx, x-1, y, l+1)
	}

	if y-1 >= 0 && mx[y-1][x] > c {
		dfs(memo, mx, x, y-1, l+1)
	}

	if h > y+1 && mx[y+1][x] > c {
		dfs(memo, mx, x, y+1, l+1)
	}

	memo.Store(x, y, l)
}

func main() {
	mx := make([][]int, 0, 200)
	for i := 0; i < 200; i++ {
		row := make([]int, 0, 200)
		for j := 0; j < 200; j++ {
			row = append(row, rand.Int())
		}
		mx = append(mx, row)
	}
	lip := longestIncreasingPath(mx)

	fmt.Println(lip)
}
