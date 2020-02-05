package bst

import (
	"github.com/golang-collections/collections/stack"
)

type Node struct {
	data        int
	left, right *Node
}

func (n *Node) New(value int) {
	n.data = value
}

type BST struct {
	root *Node
}

func (b *BST) Root() *Node {
	return b.root
}

func (b *BST) IsEmpty() bool {
	return b.root == nil
}

func (b *BST) Size() int {
	current := b.root
	size := 0
	s := stack.New()
	for s.Len() != 0 || current != nil {
		if current != nil {
			s.Push(current)
			current = current.left
		} else {
			size++
			current = s.Pop().(*Node)
			current = current.right
		}
	}
	return size
}

func (b *BST) Clear() {
	b.root = nil
}
