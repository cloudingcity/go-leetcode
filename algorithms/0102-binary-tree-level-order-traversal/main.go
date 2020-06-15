package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var (
		res   [][]int
		queue []*TreeNode
	)
	queue = append(queue, root)
	for len(queue) > 0 {
		var vals []int
		lenQueue := len(queue)
		for i := 0; i < lenQueue; i++ {
			node := queue[i]
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[lenQueue:]
		res = append(res, vals)
	}
	return res
}
