package common

import (
	"container/list"
	"fmt"
)

// Stack implements a stack data structure.
type Stack struct {
	elements *list.List
}

// NewStack initializes and returns a new stack.
func NewStack() *Stack {
	l := list.New()
	return &Stack{elements: l}
}

// Empty checks if the stack is empty.
func (q *Stack) Empty() bool {
	return q.Size() == 0
}

// Size returns the total number of elements in the stack.
func (q *Stack) Size() int {
	return q.elements.Len()
}

// Top returns the element on top of the stack.
func (q *Stack) Top() interface{} {
	return q.elements.Back().Value
}

// Push adds the element on top of the stack.
func (q *Stack) Push(v interface{}) {
	_ = q.elements.PushBack(v)
}

// Pop removes and returns the element on top of the stack.
func (q *Stack) Pop() interface{} {
	return q.elements.Remove(q.elements.Back())
}

// Print prints all elements in the stack.
func (q *Stack) Print() {
	for e := q.elements.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
