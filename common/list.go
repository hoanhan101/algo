package common

import (
	"container/list"
)

// List implements a linked list.
type List struct {
	elements *list.List
}

// NewList returns a new list.
func NewList() *List {
	return &List{elements: list.New()}
}

// Back returns the last element of the list.
func (l *List) Back() interface{} {
	return l.elements.Back().Value
}

// Front returns the first element of the list.
func (l *List) Front() interface{} {
	return l.elements.Front().Value
}

// Len returns the number of elements of the list.
func (l *List) Len() interface{} {
	return l.elements.Len()
}

// InsertAfter inserts an element after another.
func (l *List) InsertAfter(v, mark interface{}) {
	// TODO
}

// InsertBefore inserts an element before another.
func (l *List) InsertBefore(v, mark interface{}) {
	// TODO
}

// MoveAfter moves an element after another.
func (l *List) MoveAfter(v, mark interface{}) {
	// TODO
}

// MoveBefore moves an element before another.
func (l *List) MoveBefore(v, mark interface{}) {
}

// MoveBack moves an element to the back of the list.
func (l *List) MoveBack(v interface{}) {
	// TODO
}

// MoveFront moves an element to the front of the list.
func (l *List) MoveFront(v interface{}) {
	// TODO
}

// PushBack inserts an element at the back of the list.
func (l *List) PushBack(v interface{}) interface{} {
	return l.elements.PushBack(v).Value
}

// PushFront inserts an element at the front of the list.
func (l *List) PushFront(v interface{}) interface{} {
	return l.elements.PushFront(v).Value
}

// RemoveFront removes an element at the front of the list.
func (l *List) RemoveFront() interface{} {
	return l.elements.Remove(l.elements.Front())
}

// RemoveBack removes an element at the back of the list.
func (l *List) RemoveBack() interface{} {
	return l.elements.Remove(l.elements.Back())
}

// Remove removes an element from the list.
// func (l *List) Remove(v interface{}) interface{} {
// }

// Slice returns a slice of all elements in the list.
func (l *List) Slice() []interface{} {
	out := []interface{}{}

	for e := l.elements.Front(); e != nil; e = e.Next() {
		out = append(out, e.Value)
	}

	return out
}
