package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	rootIdx := 0
	for inorder[rootIdx] != rootVal {
		rootIdx++
	}
	left := buildTree(preorder[1:1+rootIdx], inorder[:rootIdx])
	right := buildTree(preorder[1+rootIdx:], inorder[rootIdx+1:])

	return &TreeNode{
		Val:   rootVal,
		Left:  left,
		Right: right,
	}
}
