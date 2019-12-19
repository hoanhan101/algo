/*
Problem:
- Given the head of a singly linked list, write a function to reorder it
  such that nodes from the second half are inserted alternately to the
  nodes from the first half in reverse order.

Approach:
- Similar to palindrome linked list problem, can also use a trick to
  reverse the second half and rearrange them in the required order
  using fast and slow pointers.

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestReorder(t *testing.T) {
	t1 := common.NewListNode(1)
	t1.AddNext(1)
	t1.AddNext(1)
	t1.AddNext(1)
	t1.AddNext(1)

	t2 := common.NewListNode(1)
	t2.AddNext(2)
	t2.AddNext(3)
	t2.AddNext(4)
	t2.AddNext(5)

	t3 := common.NewListNode(1)
	t3.AddNext(2)
	t3.AddNext(3)
	t3.AddNext(4)
	t3.AddNext(5)
	t3.AddNext(6)

	tests := []struct {
		in       *common.ListNode
		expected []int
	}{
		{t1, []int{1, 1, 1, 1, 1}},
		{t2, []int{1, 5, 2, 4, 3}},
		{t3, []int{1, 6, 2, 5, 3, 4}},
	}

	for _, tt := range tests {
		reorder(tt.in)
		common.Equal(
			t,
			tt.expected,
			common.LinkedListToSlice(tt.in),
		)
	}
}

func reorder(head *common.ListNode) {
	slow, fast := head, head

	// find the middle of the linked list.
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

	}

	// reverse a half.
	headSecondHalf := reverse(slow)
	headFirstHalf := head

	// compare the first and second half.
	for headFirstHalf != nil && headSecondHalf != nil {
		tmp := headFirstHalf.Next
		headFirstHalf.Next = headSecondHalf
		headFirstHalf = tmp

		tmp = headSecondHalf.Next
		headSecondHalf.Next = headFirstHalf
		headSecondHalf = tmp
	}

	// nil the last node's next pointer.
	if headFirstHalf != nil {
		headFirstHalf.Next = nil
	}
}
