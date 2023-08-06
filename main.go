package main

import (
	"fmt"
	"kimseogyu/tries/trie"
)

func main() {
	trie := trie.NewRadixTrie()

	trie.Insert("apple", 42)
	trie.Insert("app", 23)
	trie.Insert("banana", 56)

	if val, found := trie.Search("apple"); found {
		fmt.Println("Value for 'apple':", val)
	}

	if val, found := trie.Search("app"); found {
		fmt.Println("Value for 'app':", val)
	}

	if val, found := trie.Search("banana"); found {
		fmt.Println("Value for 'banana':", val)
	}

	if _, found := trie.Search("orange"); !found {
		fmt.Println("Key 'orange' not found.")
	}
}
