/*
Problem:
- Determine if a singly-linked list has a cycle.

Approach:
- Keep two pointers starting at the first node such that: every time one moves
  one node ahead, the other moves 2 nodes ahead.
- If the linked list has a cycle, the faster one will catch up with the slow
  one. Otherwise, the faster one will each the end.

Solution:
- Keep two pointers, a slow one and a fast one, at the beginning.
- Traverse the list and move the fast one 2 nodes once the slow one move a node.
- If the fast one catches the slow one, there exists a cycle.

Cost:
- O(n) time and O(1) space.
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestContainCycle(t *testing.T) {
	// define tests input.
	t1 := common.NewListNode(1)
	t1.AddNext(2)
	t1.AddNext(3)

	t2 := common.NewListNode(1)
	t2.AddNext(2)
	t2.AddNext(3)
	t2.Next.Next.Next = t2

	// define tests output.
	tests := []struct {
		in       *common.ListNode
		expected bool
	}{
		{t1, false},
		{t2, true},
	}

	for _, tt := range tests {
		common.Equal(t, tt.expected, containCycle(tt.in))
	}
}

func containCycle(node *common.ListNode) bool {
	// keep two pointers at the beginning.
	slow := node
	fast := node

	// traverse until it hit the end of the list.
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// if the fast pointer catches the slow one, there exists a cycle.
		if fast.Value == slow.Value {
			return true
		}
	}

	return false
}
