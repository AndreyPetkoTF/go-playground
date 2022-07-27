package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
	}

	head.Next.Next.Next = head

	fmt.Println(hasCycle(head))
}

func hasCycle(head *ListNode) bool {
	seen := make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}

		seen[head] = struct{}{}
		head = head.Next
	}

	return false
}

//func contains()
