package main

import (
	"container/heap"
	"fmt"
)

func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := NewMaxHeap(len(heights))

	for i := 0; i < len(heights)-1; i++ {
		delta := heights[i+1] - heights[i]
		if delta <= 0 {
			continue
		}

		heap.Push(h, delta)
		bricks -= delta

		if bricks < 0 {
			bricks += heap.Pop(h).(int)
			if ladders > 0 {
				ladders--
			} else {
				return i
			}
		}
	}

	return len(heights) - 1
}

func main() {
	idx := furthestBuilding([]int{14, 3, 19, 3}, 17, 0)
	fmt.Println(idx)
}
