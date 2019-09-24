// Problem:
// Determine if a singly-linked list has a cycle.
//
// Solution:
// Keep two pointers starting at the first node such that: every time one moves
// one node ahead, the other moves 2 nodes ahead.
// If the linked list has a cycle, the faster one will catch up with the slow
// one.
// Otherwise, the faster one will each the end.
//
// Cost:
// O(n) time and O(1) space.

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestContainCycle(t *testing.T) {
	// define tests input.
	t1 := &LinkedList{1, nil}
	t1.next = &LinkedList{2, nil}
	t1.next.next = &LinkedList{3, nil}

	t2 := &LinkedList{1, nil}
	t2.next = &LinkedList{2, nil}
	t2.next.next = &LinkedList{3, nil}
	t2.next.next.next = t2

	// define tests output.
	tests := []struct {
		in       *LinkedList
		expected bool
	}{
		{t1, false},
		{t2, true},
	}

	for _, tt := range tests {
		common.Equal(t, tt.expected, containCycle(tt.in))
	}
}

func containCycle(node *LinkedList) bool {
	// keep two pointers at the beginning.
	slow := node
	fast := node

	// traverse until it hit the end of the list.
	for fast != nil && fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next

		// if the fast pointer catches the slow one, there exists a cycle.
		if fast.value == slow.value {
			return true
		}
	}

	return false
}
