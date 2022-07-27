package main

import "fmt"

func NewQueue() Queue {
	return Queue{}
}

type Queue struct {
	items []int
}

func (q *Queue) Pop() int {
	item := q.items[0]

	q.items = q.items[1:]

	return item
}

func (q *Queue) Push(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) isEmpty() bool {
	return q.Size() == 0
}

func main() {
	queue := NewQueue()

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue.Size())

	item := queue.Pop()
	fmt.Println(item)

}

