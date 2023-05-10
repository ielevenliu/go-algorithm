package binarytree

import "log"

func PreTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	log.Println(root.Value)
	PreTraversal(root.Left)
	PreTraversal(root.Right)
}

func InTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	InTraversal(root.Left)
	log.Println(root.Value)
	InTraversal(root.Right)
}

func PostTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	PostTraversal(root.Left)
	PostTraversal(root.Right)
	log.Println(root.Value)
}
