package common

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	h := NewMinHeap()

	Equal(t, 0, h.Len())

	h.Push(46)
	h.Push(86)
	h.Push(6)
	h.Push(66)
	h.Push(96)
	h.Push(16)

	Equal(t, 6, h.Len())

	Equal(t, 6, h.Peek())
	Equal(t, 6, h.Pop())
	Equal(t, 16, h.Peek())
	Equal(t, 16, h.Pop())
	Equal(t, 46, h.Peek())
	Equal(t, 46, h.Pop())
	Equal(t, 66, h.Peek())
	Equal(t, 66, h.Pop())
	Equal(t, 86, h.Peek())
	Equal(t, 86, h.Pop())
	Equal(t, 96, h.Peek())
	Equal(t, 96, h.Pop())

	Equal(t, 0, h.Len())
}

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap()

	Equal(t, 0, h.Len())

	h.Push(46)
	h.Push(86)
	h.Push(6)
	h.Push(66)
	h.Push(96)
	h.Push(16)

	Equal(t, 6, h.Len())

	Equal(t, 96, h.Peek())
	Equal(t, 96, h.Pop())
	Equal(t, 86, h.Peek())
	Equal(t, 86, h.Pop())
	Equal(t, 66, h.Peek())
	Equal(t, 66, h.Pop())
	Equal(t, 46, h.Peek())
	Equal(t, 46, h.Pop())
	Equal(t, 16, h.Peek())
	Equal(t, 16, h.Pop())
	Equal(t, 6, h.Peek())
	Equal(t, 6, h.Pop())

	Equal(t, 0, h.Len())
}
