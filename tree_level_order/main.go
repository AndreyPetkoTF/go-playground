package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(levelOrder(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}))
}

func levelOrder(root *TreeNode) [][]int {
	//res := make(map[int][]int)
	res := [][]int{}

	level := 0
	maxLevel := 0
	calculateLevel(root, res, level, &maxLevel)

	fmt.Println(res)
	//var resSlice [][]int
	//for i := 0; i < maxLevel; i++ {
	//	resSlice = append(resSlice, res[i])
	//}

	return res
}

func calculateLevel(root *TreeNode, res [][]int, level int, maxLevel *int) {
	if root == nil {
		return
	}

	if len(res) < level+1 {
		res = append(res, []int{})
	}

	res[level] = append(res[level], root.Val)

	level++

	if level > *maxLevel {
		*maxLevel = level
	}
	calculateLevel(root.Left, res, level, maxLevel)
	calculateLevel(root.Right, res, level, maxLevel)
}
