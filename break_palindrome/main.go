package main

import (
	"fmt"
)

func breakPalindrome(palindrome string) string {
	chars := []byte(palindrome)
	sig := [26]int{}
	for _, char := range chars {
		sig[char-'a'] ^= 1
	}
	fmt.Println(sig)
	return ""
}

func main() {
	breakPalindrome("abccba")
}