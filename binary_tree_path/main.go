package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := binaryTreePath(&TreeNode{1, &TreeNode{2, &TreeNode{5, nil, nil}, nil}, &TreeNode{3, nil, nil}})

	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePath(root *TreeNode) []string {
	var res []string
	findPath(root, strconv.Itoa(root.Val), &res)

	return res
}

func findPath(root *TreeNode, path string, paths *[]string) {
	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, path)
		return
	}

	if root.Left != nil {
		findPath(root.Left, path+"->"+strconv.Itoa(root.Left.Val), paths)
	}

	if root.Right != nil {
		findPath(root.Left, path+"->"+strconv.Itoa(root.Right.Val), paths)
	}
}
