// Problem:
// Find the kth to last node in a linked list.
//
// Example:
// Input: list = 1 -> 2 -> 3 -> 4 -> 5 -> 6, k = 2
// Output: 5, because 5 is the 2nd to the last node (6)
//
// Solution:
// Use two pointers such that one starts at the beginning and the other one
// starts at k distance apart.
// Walk both at the same speed toward the end.
// When one hits the tail, the other one is on the target node.
//
// Cost:
// O(n) time and O(1) space.

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestKthToLast(t *testing.T) {
	// define tests input.

	t1 := &LinkedList{1, nil}

	t2 := &LinkedList{1, nil}
	t2.next = &LinkedList{2, nil}

	t3 := &LinkedList{1, nil}
	t3.next = &LinkedList{2, nil}
	t3.next.next = &LinkedList{3, nil}

	t4 := &LinkedList{1, nil}
	t4.next = &LinkedList{2, nil}
	t4.next.next = &LinkedList{3, nil}
	t4.next.next.next = &LinkedList{4, nil}
	t4.next.next.next.next = &LinkedList{5, nil}
	t4.next.next.next.next.next = &LinkedList{6, nil}

	// define tests output.
	tests := []struct {
		in1      *LinkedList
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
		common.Equal(t, tt.expected, node.value)
	}
}

func kthToLast(node *LinkedList, k int) *LinkedList {
	// start both node in the beginning.
	left, right := node, node

	// move the right one to the kth node.
	for i := 0; i < k-1; i++ {
		right = right.next
	}

	// move both pointers until the right one hits the end.
	for right.next != nil {
		left = left.next
		right = right.next
	}

	return left
}
