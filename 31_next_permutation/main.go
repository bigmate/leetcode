package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	if len(nums) <= 0 {
		return
	}

	l, r := len(nums)-2, len(nums)-1

	for l >= 0 {
		if nums[l] >= nums[r] {
			l--
			r--
			continue
		}

		sort.Ints(nums[r:])
		for i := r; i < len(nums); i++ {
			if nums[l] < nums[i] {
				nums[l], nums[i] = nums[i], nums[l]
				return
			}
		}
	}

	sort.Ints(nums)
}

func main() {
	perm := []int{1, 1}
	nextPermutation(perm)
	fmt.Println(perm)
}
