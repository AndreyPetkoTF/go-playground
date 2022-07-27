package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//res := addTwoNumbers(
	//	&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
	//	&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
	//)

	res := addTwoIterative(
		&ListNode{9, &ListNode{9, &ListNode{9, nil}}},
		&ListNode{9, nil},
	)

	fmt.Println(res)
}

func addTwoIterative(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	head := node
	addOne := false

	for l1 != nil || l2 != nil {
		sum := 0

		if l1 != nil {
			sum += l1.Val
		}

		if l2 != nil {
			sum += l2.Val
		}

		if addOne {
			sum++
		}

		newVal := sum % 10

		if sum >= 10 {
			addOne = true
		} else {
			addOne = false
		}

		node.Next = &ListNode{newVal, nil}
		node = node.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if addOne {
		node.Next = &ListNode{1, nil}
	}

	return head.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return add(l1, l2, false)
}

func add(l1 *ListNode, l2 *ListNode, addOne bool) *ListNode {
	if l1 == nil && l2 == nil {
		if addOne {
			return &ListNode{Val: 1}
		}
		return nil
	}

	l1Val := 0
	if l1 != nil {
		l1Val = l1.Val
	}

	l2Val := 0
	if l2 != nil {
		l2Val = l2.Val
	}

	sum := l1Val + l2Val
	if addOne {
		sum++
	}
	newVal := sum % 10

	if sum >= 10 {
		addOne = true
	} else {
		addOne = false
	}

	node := &ListNode{Val: newVal}

	var l1Next *ListNode
	if l1 != nil {
		l1Next = l1.Next
	}

	var l2Next *ListNode
	if l2 != nil {
		l2Next = l2.Next
	}

	node.Next = add(l1Next, l2Next, addOne)

	return node
}
