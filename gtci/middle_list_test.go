/*
Problem:
- Given the head of a singly linked list, write a function to return the
  middle value.

Approach:
- Have a slow pointer move one step at a time while the fast one move
  2 steps at a time.
- Once the fast one reaches the end, the slow is in the middle.

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFindMid(t *testing.T) {
	// define tests input.
	t1 := common.NewListNode(1)
	t1.AddNext(2)
	t1.AddNext(3)
	t1.AddNext(4)
	t1.AddNext(5)
	t1.AddNext(6)

	t2 := common.NewListNode(1)
	t2.AddNext(2)
	t2.AddNext(3)
	t2.AddNext(4)
	t2.AddNext(5)

	// define tests output.
	tests := []struct {
		in       *common.ListNode
		expected int
	}{
		{t1, 4},
		{t2, 3},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			findMid(tt.in),
		)
	}
}

func findMid(node *common.ListNode) int {
	// keep two pointers at the beginning.
	slow := node
	fast := node

	// traverse until it hit the end of the list.
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow.Value
}
