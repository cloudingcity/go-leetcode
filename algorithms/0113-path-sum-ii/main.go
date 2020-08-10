package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	backtrack(root, sum, nil, &res)
	return res
}

func backtrack(root *TreeNode, sum int, track []int, res *[][]int) {
	if root == nil {
		return
	}
	track = append(track, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == sum {
		ans := make([]int, len(track))
		copy(ans, track)
		*res = append(*res, ans)
	} else {
		backtrack(root.Left, sum-root.Val, track, res)
		backtrack(root.Right, sum-root.Val, track, res)
	}
	track = track[:len(track)-1]
}
