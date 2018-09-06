package bst

import (
	"fmt"
)

type node struct {
	data int
	left *node
	right *node
}

type Tree struct {
	root *node
}

func (tree *Tree) isEmpty() bool {
	return tree.root == nil
}

func (tree *Tree) InsertValue(value int) {
	if tree.isEmpty() {
		tree.root = &node{data: value}
	} else {
		insert(tree.root, value)
	}
}

func insert(curr *node, value int) *node {
	if curr == nil {
		return &node{data: value}
	}

	if curr.data >= value  {
		curr.left = insert(curr.left, value)
	} else {
		curr.right = insert(curr.right, value)
	}
	return curr
}

func (tree *Tree) PrintTreeInOrder() {
	if tree.isEmpty() { fmt.Println("what???") }
	printTreeInOrder(tree.root)
	fmt.Println()
}

func printTreeInOrder(curr *node) {
	if curr.left != nil {
		printTreeInOrder(curr.left)
	}
	fmt.Print(curr.data, " ")
	if curr.right != nil {
		printTreeInOrder(curr.right)
	}
}

func (tree *Tree) PrintTreePreOrder() {
	if tree.isEmpty() { fmt.Println("what???") }
	printTreePreOrder(tree.root)
	fmt.Println()
}

func printTreePreOrder(curr *node) {
	fmt.Print(curr.data, " ")
	if curr.left != nil {
		printTreePreOrder(curr.left)
	}
	if curr.right != nil {
		printTreePreOrder(curr.right)
	}
}

