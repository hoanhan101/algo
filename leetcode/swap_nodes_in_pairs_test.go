/*
Problem:
- Given a linked list, swap every two adjacent nodes and return its head.

Assumption:
- If the length of the linked list is odd, the last node should not be swapped.
- The solution should use constant space.

Example:
- Input: 1 -> 3-> 5 -> 2 -> 4-> 6
  Output: 3 -> 1-> 2 -> 5 -> 6 -> 4
- Input: 1 -> 3-> 5 -> 2 -> 4
  Output: 3 -> 1-> 2 -> 5 -> 4

Approach:
- Traverse the list and swap the nodes pairwise by adjusting where it's pointing next.

Solution:
- Initialize dummy heads to keep track of the head nodes.
- Traverse the list while the current node and its next node is not nil.
- Swap the pairs by adjusting where it's pointing next.
- Cache the current node before advancing it to the next's next node.

Cost:
- O(n) time, O(1) space where n is the length of a linked list.
*/

package leetcode

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestSwapPairs(t *testing.T) {
	t11 := &common.ListNode{1, nil}
	t11.Next = &common.ListNode{2, nil}
	t11.Next.Next = &common.ListNode{3, nil}
	t11.Next.Next.Next = &common.ListNode{4, nil}
	t11.Next.Next.Next.Next = &common.ListNode{5, nil}
	t11.Next.Next.Next.Next.Next = &common.ListNode{6, nil}

	t12 := &common.ListNode{2, nil}
	t12.Next = &common.ListNode{1, nil}
	t12.Next.Next = &common.ListNode{4, nil}
	t12.Next.Next.Next = &common.ListNode{3, nil}
	t12.Next.Next.Next.Next = &common.ListNode{6, nil}
	t12.Next.Next.Next.Next.Next = &common.ListNode{5, nil}

	tests := []struct {
		in       *common.ListNode
		expected *common.ListNode
	}{
		{t11, t12},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			swapPairs(tt.in),
		)
	}
}

func swapPairs(l *common.ListNode) *common.ListNode {
	// dummy holds the head node.
	dummy := &common.ListNode{}
	dummy.Next = l

	// current holds the current node where prev holds the previous node.
	current := l
	prev := dummy

	for current != nil && current.Next != nil {
		next := current.Next
		nextNext := current.Next.Next

		prev.Next = next

		// swap the pairs by adjusting where it's pointing next.
		next.Next = current
		current.Next = nextNext

		// cache the current node before advancing it to the next's next node.
		prev = current
		current = nextNext
	}

	return dummy.Next
}
