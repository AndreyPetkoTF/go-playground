package main

import "fmt"

func main() {
	res := mergeTwoLists(
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	)

	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	result := &ListNode{}
	tail := result

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			tail.Next = list2
			list2 = list2.Next
		} else {
			tail.Next = list1
			list1 = list1.Next
		}

		tail = tail.Next
	}

	if list1 == nil {
		tail.Next = list2
	}

	if list2 == nil {
		tail.Next = list1
	}

	return result.Next
}

func AddItem(item *ListNode, val int) *ListNode {
	item.Val = val
	item.Next = &ListNode{}

	return item.Next
}
