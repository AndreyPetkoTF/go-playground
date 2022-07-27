package main

import (
	"fmt"
)

func main() {
	queue := NewMyQueue()

	queue.Push(1)
	queue.Push(2)

	fmt.Println(queue.Peek()) // 1
	fmt.Println(queue.Pop())  // 1

	fmt.Println(queue.Empty()) // false
}

/**
queue
1 2 3


stack1
1 2 3 -> push 3 times

pop:
stack1: {}
stack2: {3 2 1} -> pop -> return 1

peek: front (last added) or stack2 first item (1) - because it is adlready reversed stack -> queue



stack1 -> cache for items
stack2 -> already reverse queue

On each

*/

var front int

type MyQueue struct {
	stack1 *Stack
	stack2 *Stack
}

func NewMyQueue() MyQueue {
	return MyQueue{NewStack(), NewStack()}
}

func (q *MyQueue) Push(x int) {
	if q.stack1.IsEmpty() {
		front = x
	}

	q.stack1.Push(x)
}

func (q *MyQueue) Pop() int {
	if q.stack2.IsEmpty() {
		for !q.stack1.IsEmpty() {
			q.stack2.Push(q.stack1.Pop())
		}
	}

	return q.stack2.Pop()
}

func (q *MyQueue) Peek() int {
	if !q.stack2.IsEmpty() {
		return q.stack2.Top()
	}

	return front
}

func (q *MyQueue) Empty() bool {
	return q.stack1.IsEmpty() && q.stack2.IsEmpty()
}

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(x int) {
	s.items = append(s.items, x)
}

func (s *Stack) Pop() int {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item
}

func (s *Stack) Top() int {
	return s.items[len(s.items)-1]
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

/**
queue
1 2 3
pop -> 1

stack
1 2 3
pop -> 3


stack1
stack2

1 2 3


**/
