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
	root := GenBinaryTreeWithParent()
	GetSuccessorTreeNode(root.Left)
}

func TestGetPrecursorTreeNode(t *testing.T) {
	root := GenBinaryTreeWithParent()
	GetPrecursorTreeNode(root.Right.Left)
}

func TestSerializePreNonRecursive(t *testing.T) {
	root := GenBinaryTreeWithStr()
	ans := SerializePreNonRecursive(root)
	t.Logf("%s", ans)
}
