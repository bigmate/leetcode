package main

import (
	"fmt"
)

func hasGroupsSizeX(deck []int) bool {
	freq := make(map[int]int)

	for _, card := range deck {
		freq[card]++
	}

	div := 0

	for _, i := range freq {
		div = gcd(div, i)
	}

	return div >= 2
}

func gcd(x, y int) int {
	for y > 0 {
		x, y = y, x%y
	}

	return x
}

func main() {
	fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2}))
}
