package trie

type node struct {
	isKey    bool
	children map[byte]*node
	value    any
}

type RadixTrie struct {
	root *node
}

func NewRadixTrie() *RadixTrie {
	return &RadixTrie{
		root: &node{
			children: make(map[byte]*node),
		},
	}
}

func (rt *RadixTrie) Insert(key string, value any) {
	rt.root.insert(key, value)
}

func (n *node) insert(key string, value any) {
	if len(key) == 0 {
		n.isKey = true
		n.value = value
		return
	}

	if child, ok := n.children[key[0]]; ok {
		child.insert(key[1:], value)
	} else {
		newChild := &node{children: make(map[byte]*node)}
		n.children[key[0]] = newChild
		newChild.insert(key[1:], value)
	}
}

func (n *node) search(key string) (any, bool) {
	if len(key) == 0 {
		if n.isKey {
			return n.value, true
		}
		return nil, false
	}

	if child, ok := n.children[key[0]]; ok {
		return child.search(key[1:])
	}

	return nil, false
}

func (rt *RadixTrie) Search(key string) (any, bool) {
	return rt.root.search(key)
}
