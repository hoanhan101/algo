/*
Problem:
- Find the kth to last node in a linked list.

Example:
- Input: list = 1 -> 2 -> 3 -> 4 -> 5 -> 6, k = 2
  Output: 5, because 5 is the 2nd to the last node (6)

Approach:
- Use two pointers such that one starts at the beginning and the other one
  starts at k distance apart.
- Walk both at the same speed toward the end.
- When one hits the tail, the other one is on the target node.

Solution:
- Start both nodes, a left one and a right one, at the beginning.
- Move the right one to the kth node.
- Move both of them until the right one hits the end.
- Return the left one.

Cost:
- O(n) time and O(1) space.
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestKthToLast(t *testing.T) {
	// define tests input.

	t1 := common.NewListNode(1)

	t2 := common.NewListNode(1)
	t2.AddNext(2)

	t3 := common.NewListNode(1)
	t3.AddNext(2)
	t3.AddNext(3)

	t4 := common.NewListNode(1)
	for i := 2; i <= 6; i++ {
		t4.AddNext(i)
	}

	// define tests output.
	tests := []struct {
		in1      *common.ListNode
		in2      int
		expected int
	}{
		{t1, 1, 1},
		{t2, 1, 2},
		{t2, 2, 1},
		{t3, 1, 3},
		{t3, 2, 2},
		{t3, 3, 1},
		{t4, 1, 6},
		{t4, 2, 5},
		{t4, 3, 4},
		{t4, 4, 3},
		{t4, 5, 2},
		{t4, 6, 1},
	}

	for _, tt := range tests {
		node := kthToLast(tt.in1, tt.in2)
		common.Equal(t, tt.expected, node.Value)
	}
}

func kthToLast(node *common.ListNode, k int) *common.ListNode {
	// start both node in the beginning.
	left, right := node, node

	// move the right one to the kth node.
	for i := 0; i < k-1; i++ {
		right = right.Next
	}

	// move both pointers until the right one hits the end.
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}

	return left
}
