package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	res := sortedArrayToBST([]int{-10, -3, 0, 5, 9})

	fmt.Println(res)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	centralIndex := len(nums) / 2

	node := &TreeNode{}

	node.Val = nums[centralIndex]

	node.Left = sortedArrayToBST(nums[:centralIndex])
	node.Right = sortedArrayToBST(nums[centralIndex+1:])

	return node
}
