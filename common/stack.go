package common

import (
	"container/list"
	"fmt"
)

type Stack struct {
	elements *list.List
}

func NewStack() *Stack {
	l := list.New()
	return &Stack{elements: l}
}

func (q *Stack) Empty() bool {
	return q.Size() == 0
}

func (q *Stack) Size() int {
	return q.elements.Len()
}

func (q *Stack) Top() interface{} {
	return q.elements.Back().Value
}

func (q *Stack) Push(v interface{}) {
	_ = q.elements.PushBack(v)
}

func (q *Stack) Pop() interface{} {
	return q.elements.Remove(q.elements.Back())
}

func (q *Stack) Print() {
	for e := q.elements.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
