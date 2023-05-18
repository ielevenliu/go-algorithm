package binarytree

import "log"

func PreRecursiveTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	log.Println(root.Value)
	PreRecursiveTraversal(root.Left)
	PreRecursiveTraversal(root.Right)
}

func InRecursiveTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	InRecursiveTraversal(root.Left)
	log.Println(root.Value)
	InRecursiveTraversal(root.Right)
}

func PostRecursiveTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	PostRecursiveTraversal(root.Left)
	PostRecursiveTraversal(root.Right)
	log.Println(root.Value)
}

func DeserializePreRecursive(serials []string, index *int) *TreeNodeWithStr {
	if len(serials) <= *index || serials[*index] == "#" {
		return nil
	}
	root := &TreeNodeWithStr{Value: serials[*index]}
	*index += 1
	root.Left = DeserializePreRecursive(serials, index)
	root.Right = DeserializePreRecursive(serials, index)
	return root
}
