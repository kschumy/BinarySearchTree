package main

import "BinarySearchTree/src"

func main() {
	foo := bst.Tree{}

	for i := 0; i < 10; i++ {
		foo.InsertValue(i)
	}

	foo.PrintTreeInOrder()
	foo.PrintTreePreOrder()
}