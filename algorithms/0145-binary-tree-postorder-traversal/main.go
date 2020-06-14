package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			res = append([]int{root.Val}, res...)
			root = root.Right
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Left
		}
	}
	return res
}

func postorderTraversalRecursive(root *TreeNode) []int {
	var res []int
	helper(root, &res)
	return res
}

func helper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		helper(root.Left, res)
	}
	if root.Right != nil {
		helper(root.Right, res)
	}
	*res = append(*res, root.Val)
}
