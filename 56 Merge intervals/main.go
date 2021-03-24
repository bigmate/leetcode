package mergeIntervals

import (
	"sort"
)

type Intervals [][]int

func (it Intervals) Len() int {
	return len(it)
}

func (it Intervals) Less(i, j int) bool {
	return it[i][0] < it[j][0]
}

func (it Intervals) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}

func merge(arr [][]int) [][]int {
	sort.Sort(Intervals(arr))
	end := 1
	res := make([][]int, 0)
	if len(arr) == 0 {
		return res
	}
	interval := arr[0]
	for interval != nil {
		if end < len(arr) && arr[end] != nil && interval[1] >= arr[end][0] {
			if interval[1] < arr[end][1] {
				interval[1] = arr[end][1]
			}
		} else {
			res = append(res, interval)
			if end < len(arr) {
				interval = arr[end]
			} else {
				interval = nil
			}
		}
		end++
	}
	return res
}
