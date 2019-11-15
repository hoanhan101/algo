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
	t1 := &common.ListNode{}

	t2 := &common.ListNode{1, nil}

	t3 := &common.ListNode{1, nil}
	t3.Next = &common.ListNode{2, nil}

	t4 := &common.ListNode{1, nil}
	t4.Next = &common.ListNode{2, nil}
	t4.Next.Next = &common.ListNode{3, nil}

	// define tests output.
	tests := []struct {
		in       *common.ListNode
		expected []int
	}{
		{t1, []int{0}},
		{t2, []int{1}},
		{t3, []int{2, 1}},
		{t4, []int{3, 2, 1}},
	}

	for _, tt := range tests {
		node := reverseLinkedList(tt.in)
		common.Equal(t, tt.expected, common.LinkedListToSlice(node))
	}
}

func reverseLinkedList(node *common.ListNode) *common.ListNode {
	current := node
	var previous, next *common.ListNode

	for current != nil {
		// copy a pointer to the next element.
		next = current.Next

		// reverse the next pointer.
		current.Next = previous

		// step forward in the list.
		previous = current
		current = next
	}

	// return the head of new linked list.
	return previous
}
