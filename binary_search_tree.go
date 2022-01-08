package main

import "fmt"

type BinaryNode struct {
	left  *BinaryNode
	data  string
	right *BinaryNode
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(data string) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data string) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func (t *BinaryNode) Inorder() {
	if t == nil {
		return
	}

	t.left.Inorder()
	fmt.Println(t.data)
	t.right.Inorder()
}

func main() {
	names_list := []string{"amaravati", "nagpur", "akola", "ratnagiri", "agra", "bhopal", "bengaluru", "chennai", "prayag", "nashik"}
	tree := &BinaryTree{}
	for _, v := range names_list {
		tree.insert(v)
	}

	tree.root.Inorder()
}
