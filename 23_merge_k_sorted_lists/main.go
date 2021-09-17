package main

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	for {
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
			doBreak = false
			break
		}
		if doBreak {
			break
		}
	}
	return head
}
