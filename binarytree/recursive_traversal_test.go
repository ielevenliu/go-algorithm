package binarytree

import "testing"

func TestPreTraversal(t *testing.T) {
	root := GenBinaryTree()
	PreTraversal(root)
}

func TestInTraversal(t *testing.T) {
	root := GenBinaryTree()
	InTraversal(root)
}

func TestPostTraversal(t *testing.T) {
	root := GenBinaryTree()
	PostTraversal(root)
}
