package main

import (
	"fmt"
)

type DifferenceFinder interface {
	UsingBitwise(s, t string) byte
	UsingSignature(s, t string) byte
	UsingSumDifference(s, t string) byte
}

type finder struct{}

func (f finder) UsingBitwise(s, t string) byte {
	var res int32

	for _, c := range s {
		res ^= c
	}
	for _, c := range t {
		res ^= c
	}

	return byte(res)
}

func (f finder) UsingSignature(s, t string) byte {
	sig := [26]int{}

	for _, char := range s {
		sig[char-'a']++
	}

	for _, char := range t {
		sig[char-'a']--
	}

	for i, count := range sig {
		if count < 0 {
			return byte(i + 'a')
		}
	}

	return 0
}

func (f finder) UsingSumDifference(s, t string) byte {
	var sum int32
	for _, char := range t {
		sum += char
	}
	for _, char := range s {
		sum -= char
	}

	return byte(sum)
}

func main() {
	s, t := "abcd", "beacd"
	find := finder{}
	fmt.Println(find.UsingSumDifference(s, t))
	fmt.Println(find.UsingSignature(s, t))
	fmt.Println(find.UsingBitwise(s, t))
}
