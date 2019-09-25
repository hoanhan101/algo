// Problem:
// Reverse a linked list in-place.
//
// Solution:
// Iterate through the list and point each node's next pointer to the previous
// item.
//
// Cost:
// O(n) time and O(1) space.

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestReverseLinkedList(t *testing.T) {
	// define tests input.
	t1 := &LinkedList{}

	t2 := &LinkedList{1, nil}

	t3 := &LinkedList{1, nil}
	t3.next = &LinkedList{2, nil}

	t4 := &LinkedList{1, nil}
	t4.next = &LinkedList{2, nil}
	t4.next.next = &LinkedList{3, nil}

	// define tests output.
	tests := []struct {
		in       *LinkedList
		expected []int
	}{
		{t1, []int{0}},
		{t2, []int{1}},
		{t3, []int{2, 1}},
		{t4, []int{3, 2, 1}},
	}

	for _, tt := range tests {
		p := reverseLinkedList(tt.in)
		common.Equal(t, tt.expected, linkedListToSlice(p))
	}
}

func reverseLinkedList(node *LinkedList) *LinkedList {
	current := node
	var previous, next *LinkedList

	for current != nil {
		// copy a pointer to the next element.
		next = current.next

		// reverse the next pointer.
		current.next = previous

		// step forward in the list.
		previous = current
		current = next
	}

	// return the head of new linked list.
	return previous
}
