package common

import (
	"container/list"
	"fmt"
)

// Queue implements a queue data structure.
type Queue struct {
	elements *list.List
}

// NewQueue initializes and returns a new queue.
func NewQueue() *Queue {
	l := list.New()
	return &Queue{elements: l}
}

// Empty checks if the queue is empty.
func (q *Queue) Empty() bool {
	return q.Size() == 0
}

// Size returns the total number of elements in the queue.
func (q *Queue) Size() int {
	return q.elements.Len()
}

// Front returns the element in the front of the queue.
func (q *Queue) Front() interface{} {
	return q.elements.Front().Value
}

// Back returns the element in the back of the queue.
func (q *Queue) Back() interface{} {
	return q.elements.Back().Value
}

// Push adds a element in the back of the queue.
func (q *Queue) Push(v interface{}) {
	_ = q.elements.PushBack(v)
}

// Pop removes and returns the element in the front of the queue.
func (q *Queue) Pop() interface{} {
	return q.elements.Remove(q.elements.Front())
}

// Print prints all the elements in the queue.
func (q *Queue) Print() {
	for e := q.elements.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
