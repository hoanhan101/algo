package common

import (
	"container/list"
	"fmt"
)

type Queue struct {
	elements *list.List
}

func NewQueue() *Queue {
	l := list.New()
	return &Queue{elements: l}
}

func (q *Queue) Empty() bool {
	return q.Size() == 0
}

func (q *Queue) Size() int {
	return q.elements.Len()
}

func (q *Queue) Front() interface{} {
	return q.elements.Front().Value
}

func (q *Queue) Back() interface{} {
	return q.elements.Back().Value
}

func (q *Queue) Push(v interface{}) {
	_ = q.elements.PushBack(v)
}

func (q *Queue) Pop() interface{} {
	return q.elements.Remove(q.elements.Front())
}

func (q *Queue) Print() {
	for e := q.elements.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
