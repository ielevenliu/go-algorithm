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

type TreeNodeWithParent struct {
	Value  int
	Left   *TreeNodeWithParent
	Right  *TreeNodeWithParent
	Parent *TreeNodeWithParent
}

func GenBinaryTreeWithParent() *TreeNodeWithParent {
	root := &TreeNodeWithParent{Value: 1}

	root.Left = &TreeNodeWithParent{Value: 2, Parent: root}
	root.Right = &TreeNodeWithParent{Value: 3, Parent: root}

	root.Left.Left = &TreeNodeWithParent{Value: 4, Parent: root.Left}
	root.Left.Right = &TreeNodeWithParent{Value: 5, Parent: root.Left}

	root.Right.Left = &TreeNodeWithParent{Value: 6, Parent: root.Right}
	root.Right.Right = &TreeNodeWithParent{Value: 7, Parent: root.Right}

	return root
}

type TreeNodeWithStr struct {
	Value string
	Left  *TreeNodeWithStr
	Right *TreeNodeWithStr
}

func GenBinaryTreeWithStr() *TreeNodeWithStr {
	root := &TreeNodeWithStr{Value: "1"}

	root.Left = &TreeNodeWithStr{Value: "2"}
	root.Right = &TreeNodeWithStr{Value: "3"}

	root.Left.Left = &TreeNodeWithStr{Value: "4"}
	root.Left.Right = &TreeNodeWithStr{Value: "5"}

	root.Right.Left = &TreeNodeWithStr{Value: "6"}
	root.Right.Right = &TreeNodeWithStr{Value: "7"}

	return root
}
