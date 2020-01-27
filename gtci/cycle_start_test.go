/*
Problem:
- Given the head of a singly linked list, write a function to find the
  starting node of the cycle.

Approach:
- Similar to finding a cycle in a linked list problem, can also determine
  the start of its cycle and calculate length k of the cycle.
- Have one pointer at the beginning and one at kth node of the linked list.
- Move both of them until they meet at the start.of the cycle.

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFindCycleStart(t *testing.T) {
	t1 := common.NewListNode(1)
	t1.AddNext(2)
	t1.AddNext(3)
	t1.AddNext(4)
	t1.AddNext(5)
	t1.AddNext(6)
	t1.Next.Next.Next.Next.Next.Next = t1.Next.Next

	t2 := common.NewListNode(1)
	t2.AddNext(2)
	t2.AddNext(3)
	t2.AddNext(4)
	t2.AddNext(5)
	t2.AddNext(6)
	t2.Next.Next.Next.Next.Next.Next = t2.Next.Next.Next

	t3 := common.NewListNode(1)
	t3.AddNext(2)
	t3.AddNext(3)
	t3.AddNext(4)
	t3.AddNext(5)
	t3.AddNext(6)
	t3.Next.Next.Next.Next.Next.Next = t3

	t4 := common.NewListNode(1)
	t4.AddNext(2)
	t4.AddNext(3)
	t4.AddNext(4)
	t4.AddNext(5)
	t4.AddNext(6)

	tests := []struct {
		in       *common.ListNode
		expected int
	}{
		{t1, 3},
		{t2, 4},
		{t3, 1},
		{t4, 1},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			findCycleStart(tt.in),
		)
	}
}

func findCycleStart(head *common.ListNode) int {
	length := 0

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			length = cycleLength(slow)
			break
		}
	}

	start := findStart(head, length)
	return start.Value
}

// cycleLength loops around the cycle and calculate its length.
func cycleLength(head *common.ListNode) int {
	length := 0

	current := head
	for {
		current = current.Next
		length++

		if current == head {
			break
		}
	}

	return length
}

func findStart(head *common.ListNode, k int) *common.ListNode {
	slow, fast := head, head

	// move the fast pointer k distance ahead.
	for k > 0 {
		fast = fast.Next
		k--
	}

	// increment both pointers until they meet at the start.
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
