package main

import (
	"fmt"
	"math"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(maxPathSum(&TreeNode{
		Val: -1,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:  4,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   -4,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: nil,
	})) // 42
}

func maxPathSum(root *TreeNode) int {
	goose := math.MinInt
	calculateForNode(root, &goose)

	return goose
}

func calculateForNode(node *TreeNode, goose *int) int {
	if node.Left == nil && node.Right == nil {
		if *goose < node.Val {
			*goose = node.Val
		}
		return node.Val
	}

	if node.Left != nil && node.Right != nil {
		resL := calculateForNode(node.Left, goose)
		resR := calculateForNode(node.Right, goose)

		maxValue := max([]int{resL + node.Val + resR, node.Val, resL + node.Val, resR + node.Val, resL, resR})
		if maxValue > *goose {
			*goose = maxValue
		}

		return max([]int{resL + node.Val, resR + node.Val, node.Val})
	}

	if node.Right != nil {
		resR := calculateForNode(node.Right, goose)
		maxValue := max([]int{node.Val, resR + node.Val, resR})
		if maxValue > *goose {
			*goose = maxValue
		}

		return max([]int{resR + node.Val, node.Val})
	}

	if node.Left != nil {
		resL := calculateForNode(node.Left, goose)
		maxValue := max([]int{node.Val, resL + node.Val, resL})
		if maxValue > *goose {
			*goose = maxValue
		}

		return max([]int{resL + node.Val, node.Val})
	}

	return 0
}

func max(res []int) int {
	sort.Ints(res)

	return res[len(res)-1]
}
