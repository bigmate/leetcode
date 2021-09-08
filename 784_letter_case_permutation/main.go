package main

import (
	"fmt"

	"leetcode/stack"
)

const asciiCaseMask = 0x20

func letterCasePermutation(s string) []string {
	bucket := make([]string, 0)
	st := stack.New()
	bs := []byte(s)
	generate(&bucket, st, bs, 0)
	return bucket
}

func generate(bucket *[]string, s stack.Stack, bs []byte, pos int) {
	if pos >= len(bs) {
		*bucket = append(*bucket, stackContent(s))
		return
	}
	s.Push(bs[pos])
	generate(bucket, s, bs, pos+1)
	s.Pop()
	if caseSwappable(bs[pos]) {
		s.Push(swapCase(bs[pos]))
		generate(bucket, s, bs, pos+1)
		s.Pop()
	}
}

func caseSwappable(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z')
}

func swapCase(char byte) byte {
	if caseSwappable(char) {
		return char ^ asciiCaseMask
	}
	return char
}

func stackContent(s stack.Stack) string {
	buf := make([]byte, 0, s.Size())
	iter := s.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		buf = append(buf, data.(byte))
	}
	return string(buf)
}

func main() {
	cmb := letterCasePermutation("a1b2")
	fmt.Println(cmb)
}
