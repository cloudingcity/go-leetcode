package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var (
		res   [][]int
		queue []*TreeNode
	)
	foward := true
	queue = append(queue, root)
	for len(queue) > 0 {
		var vals []int
		lenQueue := len(queue)
		for i := lenQueue - 1; i >= 0; i-- {
			node := queue[i]
			vals = append(vals, node.Val)

			if foward {
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			} else {
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
			}
		}
		res = append(res, vals)
		queue = queue[lenQueue:]
		foward = !foward
	}
	return res
}
