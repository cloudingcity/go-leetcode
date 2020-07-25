package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}

func helper(root, lower, upper *TreeNode) bool {
	if root == nil {
		return true
	}
	if lower != nil && root.Val <= lower.Val {
		return false
	}
	if upper != nil && root.Val >= upper.Val {
		return false
	}
	return helper(root.Left, lower, root) && helper(root.Right, root, upper)
}
