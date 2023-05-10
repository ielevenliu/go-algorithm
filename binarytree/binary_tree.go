package binarytree

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func GenBinaryTree() *TreeNode {
	root := &TreeNode{Value: 1}

	root.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{Value: 3}

	root.Left.Left = &TreeNode{Value: 4}
	root.Left.Right = &TreeNode{Value: 5}

	root.Right.Left = &TreeNode{Value: 6}
	root.Right.Right = &TreeNode{Value: 7}

	return root
}
