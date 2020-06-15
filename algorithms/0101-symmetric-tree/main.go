package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var queue []*TreeNode
	queue = append(queue, root, root)

	for len(queue) > 0 {
		t1, t2 := queue[0], queue[1]
		queue = queue[2:]

		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil {
			return false
		}
		if t1.Val != t2.Val {
			return false
		}
		queue = append(queue, t1.Left, t2.Right, t1.Right, t2.Left)
	}
	return true
}

func isSymmetricRecursive(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}

	return t1.Val == t2.Val &&
		isMirror(t1.Left, t2.Right) &&
		isMirror(t1.Right, t2.Left)
}
