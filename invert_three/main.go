package main

import "fmt"

func main() {
	root := &ThreeNode{
		Val: 2,
		Left: &ThreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &ThreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	res := invertThree(root)

	fmt.Println(res)
}

type ThreeNode struct {
	Val   int
	Left  *ThreeNode
	Right *ThreeNode
}

func invertThree(root *ThreeNode) *ThreeNode {
	if root == nil {
		return nil
	}

	if root.Left != nil {
		invertThree(root.Left)
	}

	if root.Right != nil {
		invertThree(root.Right)
	}

	root.Left, root.Right = root.Right, root.Left

	return root
}
