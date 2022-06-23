package main

import "fmt"

func longestPalindrome(s string) string {
	bs := []byte(s)
	sb := bs[:1]

	for i := 0; i < len(s); i++ {
		if i != 0 {
			sb = longestRange(bs, sb, i-1, i+1)
		}
		sb = longestRange(bs, sb, i, i+1)
	}

	return string(sb)
}

func longestRange(bs, sub []byte, a, b int) []byte {
	var s []byte

	for a >= 0 && b < len(bs) && bs[a] == bs[b] {
		s = bs[a : b+1]
		a--
		b++
	}

	if len(sub) < len(s) {
		return s
	}

	return sub
}

func main() {
	fmt.Println(longestPalindrome("abad"))
}
