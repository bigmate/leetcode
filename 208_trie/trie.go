package main

type Trie struct {
	root *node
}

func Constructor() Trie {
	return Trie{root: &node{}}
}

func (t *Trie) Insert(word string) {
	cur := t.root
	for _, r := range word {
		c := byte(r)
		if cur.has(c) {
			cur = cur.get(c)
		} else {
			cur = cur.add(c)
		}
	}
	cur.end = true
}

func (t *Trie) Search(word string) bool {
	cur := t.root
	for _, r := range word {
		c := byte(r)
		if !cur.has(c) {
			return false
		}
		cur = cur.get(c)
	}
	return cur.end
}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t.root
	for _, r := range prefix {
		c := byte(r)
		if !cur.has(c) {
			return false
		}
		cur = cur.get(c)
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
