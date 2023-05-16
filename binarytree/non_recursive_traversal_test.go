package binarytree

import "testing"

func TestPreNonRecursiveTraversal(t *testing.T) {
	root := GenBinaryTree()
	PreNonRecursiveTraversal(root)
}
