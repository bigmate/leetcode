package main

type node struct {
	end bool
	arr [26]*node
}

func (n *node) add(char byte) *node {
	created := &node{}
	n.arr[char-'a'] = created
	return created
}

func (n *node) get(char byte) *node {
	checkLegal(char)
	return n.arr[char-'a']
}

func (n *node) has(char byte) bool {
	checkLegal(char)
	return n.arr[char-'a'] != nil
}

func legalChar(char byte) bool {
	return char >= 'a' && char <= 'z'
}

func checkLegal(char byte) {
	if !legalChar(char) {
		panic("bad char")
	}
}