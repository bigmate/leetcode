package main

import (
	"fmt"
)

var charsMap = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	total := 1

	for i := 0; i < len(digits); i++ {
		total *= len(charsMap[digits[i]])
	}

	bs := []byte(digits)
	res := make([]string, 0, total)

	generate(&res, bs, 0)

	return res
}

func generate(bucket *[]string, digits []byte, current int, bs ...byte) {
	if len(bs) == len(digits) {
		*bucket = append(*bucket, string(bs))
		return
	}
	if current >= len(digits) {
		return
	}
	digit := digits[current]
	chars := charsMap[digit]
	for _, char := range chars {
		generate(bucket, digits, current+1, append(bs, char)...)
	}
}

/*
      1  2  3
-----[a][d][g]------
      b  e  h
	  c  f  i
*/

func main() {
	res := letterCombinations("4")
	fmt.Println(res)
}
