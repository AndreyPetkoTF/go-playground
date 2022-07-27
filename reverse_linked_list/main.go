package main

import (
	"fmt"
)

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
				Val:  3,
				Next: nil,
			},
		},
	}

	//fmt.Println(reverseList(head))
	res := reverseListIterative(head)
	fmt.Println(res)
}

/**
1 -> 2 -> 3 -> nil
3 -> 2 -> 1 -> nil

curr = 1
curr.Next, nextTemp = 2
curr.Next.Next = 3

1 -> nil

*/

func reverseListIterative(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr

		curr = nextTemp
	}

	return prev
}

/**
1 -> 2 -> 3

head -> 1
head.Next -> 2
head.Next.Next = asd -> change link for 2

2 -> 1
1 -> nil

next call
3 -> 2
3 -> nil

next call
3 -> return 3


*/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return res
}
