package main

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	sort.Slice(lists, func(i, j int) bool {
		if lists[i] == nil && lists[j] != nil {
			return true
		}
		if lists[j] == nil && lists[i] != nil {
			return false
		}
		if lists[i] == nil && lists[j] == nil {
			return false
		}
		return lists[i].Val < lists[j].Val
	})
	var head *ListNode
	var tail *ListNode
	for {
		doBreak := true
		for i, node := range lists {
			if node == nil {
				continue
			}
			if head == nil {
				head = node
				tail = node
			} else {
				tail.Next = node
				tail = node
			}
			lists[i] = lists[i].Next
			restoreOrder(lists, i)
			doBreak = false
			break
		}
		if doBreak {
			break
		}
	}
	return head
}

func restoreOrder(lists []*ListNode, idx int) {
	if lists[idx] == nil {
		return
	}
	for i := idx; i < len(lists)-1; i++ {
		if lists[i].Val > lists[i+1].Val {
			lists[i], lists[i+1] = lists[i+1], lists[i]
		}
	}
}
