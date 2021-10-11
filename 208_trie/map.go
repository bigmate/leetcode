package main

type node struct {
	end  bool
	arr  [26]*node
}

func (n *node) add(char byte) *node {
	if !legalChar(char) {
		panic("bad char")
	}
	created := &node{}
	n.arr[char-'a'] = created
	return created
}

func (n *node) get(char byte) *node {
	if !legalChar(char) {
		panic("bad char")
	}
	return n.arr[char-'a']
}

func (n *node) has(char byte) bool {
	if !legalChar(char) {
		panic("bad char")
	}
	return n.arr[char-'a'] != nil
}

func legalChar(char byte) bool {
	return char >= 'a' && char <= 'z'
}
