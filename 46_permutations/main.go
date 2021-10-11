package main

import (
	"fmt"

	"leetcode/stack"
)

func permute(nums []int) [][]int {
	var generate func(bucket *[][]int, stack stack.Stack, nums []int, counter map[int]int)
	generate = func(bucket *[][]int, stack stack.Stack, nums []int, counter map[int]int) {
		if stack.Size() == len(nums) {
			*bucket = append(*bucket, stackContent(stack))
			return
		}
		for num, count := range counter {
			if count <= 0 {
				continue
			}
			counter[num]--
			stack.Push(num)
			generate(bucket, stack, nums, counter)
			stack.Pop()
			counter[num]++
		}
	}
	counter := make(map[int]int, len(nums))
	bucketLen := 1
	for i, num := range nums {
		bucketLen *= i + 1
		counter[num]++
	}
	bucket := make([][]int, 0, bucketLen)
	generate(&bucket, stack.New(), nums, counter)
	return bucket
}

func stackContent(stack stack.Stack) []int {
	iter := stack.Iter()
	res := make([]int, 0, stack.Size())
	for data, exists := iter.Next(); exists; data, exists = iter.Next() {
		res = append(res, data.(int))
	}
	return res
}

func main() {
	fmt.Println(permute([]int{2, 2, 3}))
}
