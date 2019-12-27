package common

import (
	"testing"
)

func TestListNodeAddNext(t *testing.T) {
	l11 := NewListNode(1)
	l11.AddNext(2)

	l12 := NewListNode(1)
	l12.Next = NewListNode(2)
	Equal(t, l12, l11)
}

func TestLinkedListToSlice(t *testing.T) {
	t1 := &ListNode{1, nil}

	t2 := &ListNode{1, nil}
	t2.Next = &ListNode{2, nil}

	t3 := &ListNode{1, nil}
	t3.Next = &ListNode{2, nil}
	t3.Next.Next = &ListNode{3, nil}

	Equal(t, []int{1}, LinkedListToSlice(t1))
	Equal(t, []int{1, 2}, LinkedListToSlice(t2))
	Equal(t, []int{1, 2, 3}, LinkedListToSlice(t3))
}
