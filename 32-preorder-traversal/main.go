package main

import "fmt"

type Node struct {
	data        string
	left, right *Node
}

func (n *Node) New(value string) {
	n.data = value
}

func (n *Node) IsLeaf() bool {
	if n.left == nil && n.right == nil {
		return true
	}
	return false
}

type Tree struct {
	root *Node
}

func (t *Tree) PreOrder(node *Node) {
	if node == nil {
		return
	} else {
		fmt.Printf("%s ", node.data)
		t.PreOrder(node.left)
		t.PreOrder(node.right)
	}
}

func main() {
	tree := Tree{}
	tree.root = &Node{data: "1"}
	tree.root.left = &Node{data: "2"}
	tree.root.left.left = &Node{data: "3"}

	tree.root.left.right = &Node{data: "4"}
	tree.root.right = &Node{data: "5"}
	tree.root.right.right = &Node{data: "6"}

	tree.PreOrder(tree.root)
	fmt.Println("")
}
