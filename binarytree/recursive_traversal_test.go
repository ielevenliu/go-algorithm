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

func TestDeserializePreRecursive(t *testing.T) {
	var index int
	serials := []string{"1", "2", "4", "#", "#", "5", "#", "#", "3", "6", "#", "#", "7", "#", "#"}
	root := DeserializePreRecursive(serials, &index)
	ans := SerializePreNonRecursive(root)
	t.Logf("%s", ans)
}
