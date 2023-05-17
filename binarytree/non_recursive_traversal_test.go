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

func TestGetFollowUpTreeNode(t *testing.T) {
	root := GenBinaryTreeNodeWithParent()
	GetFollowUpTreeNode(root.Left)
}
