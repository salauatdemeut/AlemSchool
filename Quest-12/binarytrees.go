package main

import (
	"fmt"
)

type node struct {
	left, right *node
	data        int
}

func insert(root *node, elem int) *node {
	if root == nil {
		return &node{data: elem}
	}
	if elem < root.data {
		root.left = insert(root.left, elem)
	} else {
		root.right = insert(root.right, elem)
	}
	return root
}

func PrintTree(root *node) {
	if root == nil {
		return
	}
	fmt.Println(root.data)
	PrintTree(root.left)
	PrintTree(root.right)
}

func main() {
	var root *node
	root = insert(root, 4)
	root = insert(root, 2)
	root = insert(root, 6)
	root = insert(root, 3)
	root = insert(root, 5)
	root = insert(root, 7)
	root = insert(root, 1)
	PrintTree(root)
}

/*
if 4 > 2 && 4 < 6 {
	3
}
if 4 > 2 && 2 > 1 {
	1
}
if 4 < 6 && 6 < 7 {
	5
}
if 4 < 6 && 6 > 5 {
	7
}
*/
