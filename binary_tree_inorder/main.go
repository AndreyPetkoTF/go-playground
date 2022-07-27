package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(inorderTraversal(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}))
}

func inorderTraversal(root *TreeNode) []int {
	return getValues(root)
}

func getValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	left := getValues(root.Left)
	left = append(left, root.Val)
	right := getValues(root.Right)

	return append(left, right...)
}
