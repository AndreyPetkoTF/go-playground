package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(isSymmetric(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
	}))
}

func isSymmetric(root *TreeNode) bool {
	return isSameTree(root.Left, root.Right)
}

func isSameTree(q *TreeNode, b *TreeNode) bool {
	if q == nil && b == nil {
		return true
	}

	if q == nil && b != nil || q != nil && b == nil {
		return false
	}

	if q.Val != b.Val {
		return false
	}

	return isSameTree(q.Left, b.Left) && isSameTree(q.Right, b.Right)
}
