/*
Problem:
- Given the head of a singly linked list, write a function to determine
  if it contains a cycle.

Approach:
- Have a slow pointer move one step at a time while the fast one move
  2 steps at a time.
- If the linked list has a cycle, the fast pointer will catch the slow one.

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestHasCycle(t *testing.T) {
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
		common.Equal(
			t,
			tt.expected,
			hasCycle(tt.in),
		)
	}
}

func hasCycle(node *common.ListNode) bool {
	// keep two pointers at the beginning.
	slow := node
	fast := node

	// traverse until it hit the end of the list.
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// if the fast pointer catches the slow one, there exists a cycle.
		if fast.Value == slow.Value {
			return true
		}
	}

	return false
}
