package common

import (
	"testing"
)

func TestList(t *testing.T) {
	l := NewList()

	Equal(t, 0, l.Len())
	Equal(t, []interface{}{}, l.Slice())

	// note that we don't use type assertion here because Equal uses DeepEqual
	// to compare these 2 values.
	Equal(t, 6, l.PushFront(6))
	Equal(t, 1, l.Len())
	Equal(t, []interface{}{6}, l.Slice())

	Equal(t, 1, l.PushFront(1))
	Equal(t, 2, l.Len())
	Equal(t, []interface{}{1, 6}, l.Slice())

	Equal(t, 16, l.PushBack(16))
	Equal(t, 3, l.Len())
	Equal(t, []interface{}{1, 6, 16}, l.Slice())

	Equal(t, 66, l.PushBack(66))
	Equal(t, 4, l.Len())
	Equal(t, []interface{}{1, 6, 16, 66}, l.Slice())

	Equal(t, 1, l.Front())
	Equal(t, 66, l.Back())

	Equal(t, 1, l.RemoveFront())
	Equal(t, 3, l.Len())
	Equal(t, []interface{}{6, 16, 66}, l.Slice())

	Equal(t, 66, l.RemoveBack())
	Equal(t, 2, l.Len())
	Equal(t, []interface{}{6, 16}, l.Slice())
}
