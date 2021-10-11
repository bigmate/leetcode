package main

import (
	"container/list"
	"fmt"
)

type pair struct {
	key   int
	value int
}

type LRUCache struct {
	memo     map[int]*list.Element
	recent   *list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		memo:     make(map[int]*list.Element, capacity),
		recent:   list.New(),
		capacity: capacity,
	}
}

func (lru *LRUCache) Get(key int) int {
	if el, ok := lru.memo[key]; ok {
		lru.recent.MoveToFront(el)
		return el.Value.(pair).value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	el, ok := lru.memo[key]
	if ok {
		lru.recent.Remove(el)
		delete(lru.memo, key)
	}
	if len(lru.memo) >= lru.capacity {
		back := lru.recent.Back()
		delete(lru.memo, back.Value.(pair).key)
		lru.recent.Remove(back)
	}
	front := lru.recent.PushFront(pair{
		key:   key,
		value: value,
	})
	lru.memo[key] = front
}

func main() {
	lru := Constructor(2)
	input := [][]int{{2, 1}, {2, 2}, {2}, {1, 1}, {4, 1}, {2}}
	for _, ints := range input {
		if len(ints) == 2 {
			lru.Put(ints[0], ints[1])
			fmt.Print("null, ")
		} else if len(ints) == 1 {
			fmt.Printf("%v, ", lru.Get(ints[0]))
		}
	}
}
