package main

import "fmt"

func breakPalindrome(palindrome string) string {
	chars := []byte(palindrome)
	if len(chars) <= 1 {
		return ""
	}

	for i := 0; i < len(chars)/2; i++ {
		if chars[i] != 'a' {
			chars[i] = 'a'
			return string(chars)
		}
	}

	chars[len(chars)-1] = 'b'

	return string(chars)
}

func main() {
	fmt.Println(breakPalindrome("acca"))
}
