package common

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()

	Equal(t, 0, q.Size())

	q.Push(1)
	q.Push(2)
	q.Push(3)

	Equal(t, 3, q.Size())

	Equal(t, 1, q.Front())
	Equal(t, 3, q.Back())

	Equal(t, 1, q.Pop())
	Equal(t, 2, q.Pop())
	Equal(t, 3, q.Pop())

	Equal(t, 0, q.Size())
}
