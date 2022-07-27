package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(isBalanced(&TreeNode{
		1,
		&TreeNode{2,
			&TreeNode{
				Val: 3,
			},
			nil,
		},
		&TreeNode{Val: 2},
	}))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	maxL := maxDepth(root.Left)
	maxR := maxDepth(root.Right)

	diff := maxL - maxR

	return diff >= -1 && diff <= 1
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	if l > r {
		l++
		return l
	} else {
		r++
		return r
	}
}
