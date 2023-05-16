package binarytree

import (
	"log"
)

func PreNonRecursiveTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	log.Printf("Pre-non-recursive-traversal: ")
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		node := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		log.Printf("%d ", node.Value)
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
	}
}
