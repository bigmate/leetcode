package partitionLabels

import (
	"fmt"
	"sort"
)

type Pos struct {
	start int
	end   int
}

func (p *Pos) String() string {
	return fmt.Sprintf("Pos[%v, %v]", p.start, p.end)
}

type Range [26]*Pos

func (r *Range) Len() int {
	return len(r)
}

func (r *Range) Less(i, j int) bool {
	if r[i] == nil {
		return false
	}
	if r[j] == nil {
		return true
	}
	return r[i].start < r[j].start
}

func (r *Range) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r *Range) Set(char rune, i int) {
	p := char - 'a'
	if pos := r[p]; pos != nil {
		pos.end = i
	} else {
		r[p] = &Pos{
			start: i,
			end:   i,
		}
	}
}

func partitionLabels(S string) []int {
	arr := &Range{}
	for i, char := range S {
		arr.Set(char, i)
	}
	sort.Sort(arr)
	end := 1
	res := make([]int, 0)
	pos := arr[0]
	for pos != nil {
		if end < 26 && arr[end] != nil && pos.end > arr[end].start {
			if pos.end < arr[end].end {
				pos.end = arr[end].end
			}
		} else {
			res = append(res, pos.end-pos.start+1)
			if end < 26 {
				pos = arr[end]
			} else {
				pos = nil
			}
		}
		end++
	}
	return res
}
