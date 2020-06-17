package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 && len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	rootIdx := 0
	for inorder[rootIdx] != rootVal {
		rootIdx++
	}
	left := buildTree(inorder[:rootIdx], postorder[:rootIdx])
	right := buildTree(inorder[1+rootIdx:], postorder[rootIdx:len(postorder)-1])
	return &TreeNode{
		Val:   rootVal,
		Left:  left,
		Right: right,
	}
}
