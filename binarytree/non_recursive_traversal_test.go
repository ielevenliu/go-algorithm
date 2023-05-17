package binarytree

import "testing"

func TestPreNonRecursiveTraversal(t *testing.T) {
	root := GenBinaryTree()
	PreNonRecursiveTraversal(root)
}

func TestInNonRecursiveTraversal(t *testing.T) {
	root := GenBinaryTree()
	InNonRecursiveTraversal(root)
}

func TestPostNonRecursiveTraversal(t *testing.T) {
	root := GenBinaryTree()
	PostNonRecursiveTraversal(root)
}

func TestGetSuccessorTreeNode(t *testing.T) {
	root := GenBinaryTreeNodeWithParent()
	GetSuccessorTreeNode(root.Left)
}

func TestGetPrecursorTreeNode(t *testing.T) {
	root := GenBinaryTreeNodeWithParent()
	GetPrecursorTreeNode(root.Right.Left)
}
