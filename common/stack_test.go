package common

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()

	Equal(t, 0, s.Size())

	s.Push(1)
	s.Push(2)
	s.Push(3)

	Equal(t, 3, s.Size())

	Equal(t, 3, s.Top())
	Equal(t, 3, s.Pop())
	Equal(t, 2, s.Pop())
	Equal(t, 1, s.Pop())

	Equal(t, 0, s.Size())
}
