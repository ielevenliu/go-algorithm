package binarytree

import "testing"

func TestPreTraversal(t *testing.T) {
	root := GenBinaryTree()
	PreRecursiveTraversal(root)
}

func TestInTraversal(t *testing.T) {
	root := GenBinaryTree()
	InRecursiveTraversal(root)
}

func TestPostTraversal(t *testing.T) {
	root := GenBinaryTree()
	PostRecursiveTraversal(root)
}
