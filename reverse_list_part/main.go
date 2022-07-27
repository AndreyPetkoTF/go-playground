package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	res := reverseBetween(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2, 4)

	fmt.Println(res)
}

/**
1 -> 2 -> 3
*/

var leftItem *ListNode
var stop bool

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	leftItem = head
	stop = false

	recurseAndReverse(head, left, right)

	return head
}

/**

 */

func recurseAndReverse(rightItem *ListNode, left int, right int) {
	if right == 1 {
		return
	}

	rightItem = rightItem.Next

	if left > 1 {
		leftItem = leftItem.Next
	}

	recurseAndReverse(rightItem, left-1, right-1)

	if leftItem == rightItem || rightItem.Next == leftItem {
		stop = true
	}

	if !stop {
		temp := leftItem.Val
		leftItem.Val = rightItem.Val
		rightItem.Val = temp

		leftItem = leftItem.Next
	}
}
