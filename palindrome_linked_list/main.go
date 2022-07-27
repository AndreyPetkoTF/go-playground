package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	res := isPalindrome(&ListNode{Val: 1, Next: &ListNode{2, &ListNode{2, &ListNode{1, nil}}}})

	fmt.Println(res)
}

var pointer *ListNode

func isPalindrome(head *ListNode) bool {
	pointer = head

	return recursivelyCheck(head)
}

func recursivelyCheck(node *ListNode) bool {
	if node != nil {
		if !recursivelyCheck(node.Next) {
			return false
		}

		if node.Val != pointer.Val {
			return false
		}

		pointer = pointer.Next
	}

	return true
}
