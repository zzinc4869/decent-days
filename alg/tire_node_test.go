package alg

import (
	"fmt"
	"testing"
)

type TrieNode struct {
	isEnd bool
	next  map[string]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{isEnd: false, next: make(map[string]*TrieNode)}
}

func (t *TrieNode) insert(word string) {
	curr := t
	for _, ch := range word {
		this := string(ch)
		if _, ok := curr.next[this]; !ok {
			curr.next[this] = NewTrieNode()
		}
		curr = curr.next[this]
	}
	curr.isEnd = true
}

func (t *TrieNode) hasPrefix(prefix string) bool {
	curr := t
	for _, ch := range prefix {
		this := string(ch)
		if _, ok := curr.next[this]; !ok {
			return false
		}
		curr = curr.next[this]
	}
	return true
}

func (t *TrieNode) hasWord(word string) bool {
	curr := t
	for _, ch := range word {
		this := string(ch)
		if _, ok := curr.next[this]; !ok {
			return false
		}
		curr = curr.next[this]
	}
	return curr.isEnd
}

func (t *TrieNode) print() {
	printTrie("", *t)
}

func printTrie(word string, node TrieNode) {
	if node.isEnd {
		fmt.Println(word)
	}
	for k, v := range node.next {
		printTrie(word+k, *v)
	}
}

func TestTireNode(t *testing.T) {
	node := NewTrieNode()
	node.insert("cate")
	node.insert("cat")
	node.insert("cell")
	node.print()
	fmt.Println(node.hasPrefix("cat"))
	fmt.Println(node.hasPrefix("cet"))
}
