package main

type WordDictionary struct {
	root *node
}

func Constructor() WordDictionary {
	return WordDictionary{root: &node{}}
}

func (d *WordDictionary) AddWord(word string) {
	cur := d.root
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

func (d *WordDictionary) Search(word string) bool {
	return dfs(d.root, []byte(word), 0)
}

func dfs(root *node, word []byte, pos int) bool {
	if pos >= len(word) {
		return root.end
	}
	if word[pos] == '.' {
		for _, n := range root.arr {
			if n == nil {
				continue
			}
			if dfs(n, word, pos+1) {
				return true
			}
		}
	} else if root.has(word[pos]) {
		return dfs(root.get(word[pos]), word, pos+1)
	}
	return false
}
