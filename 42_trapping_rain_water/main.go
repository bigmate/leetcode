package main

import (
	"fmt"
)

func trap(heights []int) int {
	if len(heights) <= 0 {
		return 0
	}

	totalVolume := 0

	els := maxElevations(heights)

	for i, height := range heights {
		e := els[i]
		h := min(e.R, e.L)
		totalVolume += max(h-height, 0)
	}

	return totalVolume
}

type LR struct {
	L, R int
}

func maxElevations(heights []int) []LR {
	el := make([]LR, len(heights))

	maxEl := 0
	for i := 0; i < len(heights); i++ {
		el[i].L = maxEl
		maxEl = max(heights[i], maxEl)
	}

	maxEl = 0
	for i := len(heights) - 1; i >= 0; i-- {
		el[i].R = maxEl
		maxEl = max(heights[i], maxEl)
	}

	return el
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func main() {
	v := trap([]int{4, 3, 2, 1, 0, 5})
	fmt.Println(v)
}
