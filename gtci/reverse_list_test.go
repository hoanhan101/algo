/*
Problem:
- Given the head of a singly linked list, write a function to return the
  new head of the reversed linked list.

Approach:
- Iterate through the linked list and at each step, reverse the current
  node by pointing it to the previous node before moving on.

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestReverseList(t *testing.T) {
	t0 := common.NewListNode(0)

	t1 := common.NewListNode(1)

	t2 := common.NewListNode(1)
	t2.AddNext(2)

	t3 := common.NewListNode(1)
	t3.AddNext(2)
	t3.AddNext(3)

	tests := []struct {
		in       *common.ListNode
		expected []int
	}{
		{t0, []int{0}},
		{t1, []int{1}},
		{t2, []int{2, 1}},
		{t3, []int{3, 2, 1}},
	}

	for _, tt := range tests {
		h := reverseList(tt.in)
		common.Equal(
			t,
			tt.expected,
			common.LinkedListToSlice(h),
		)
	}
}

func reverseList(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current := head
	var previous, next *common.ListNode

	for current != nil {
		// temporarily store the next node.
		next = current.Next

		// reverse the current node.
		current.Next = previous

		// point the previous node to the current node.
		previous = current

		// move on to the next node
		current = next
	}

	return previous
}
