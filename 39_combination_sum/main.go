package main

import (
	"fmt"
	"sort"

	"leetcode/stack"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	buf := make([][]int, 0)
	generate(&buf, stack.New(), candidates, target, 0)
	return buf
}

func generate(buf *[][]int, stack stack.Stack, candidates []int, target, pos int) {
	if target == 0 {
		*buf = append(*buf, stackContent(stack))
		return
	}
	for pos < len(candidates) {
		if target-candidates[pos] >= 0 {
			stack.Push(candidates[pos])
			generate(buf, stack, candidates, target-candidates[pos], pos)
			stack.Pop()
			pos++
		} else {
			break
		}
	}
}

func stackContent(stack stack.Stack) []int {
	res := make([]int, 0, stack.Size())
	iter := stack.Iter()
	for data, exists := iter.Next(); exists; data, exists = iter.Next() {
		res = append(res, data.(int))
	}
	return res
}

func main() {
	res := combinationSum([]int{1}, 2)
	fmt.Println(res)
}
