package main

import "fmt"

func main() {
	stack := NewStack()

	stack.Push(1)
	stack.Push(2)

	item := stack.Pop()

	fmt.Println(item)
}

type Stack struct {
	items []int
}

func NewStack() Stack {
	return Stack{}
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

func (s *Stack) isEmpty() bool {
	return s.Size() == 0
}
