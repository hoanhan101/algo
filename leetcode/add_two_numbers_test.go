/*
Problem:
- Given two linked lists representing two non-negative number, add them together
  and return it as a linked list.

Assumption:
- The digits are stored in reverse order.
- Each node contains a single digit.

Example:
- Input: (1 -> 6 -> 4) + (2 -> 4-> 1)
  Output: (3 -> 0 -> 6)

Approach:
- Traverse both lists and keep track of the sum and carry for each
  digit.

Solution:
- Initialize dummy heads to keep track of the head nodes.
- While either list is not empty, calculate the sum of corresponding
  digits, update it carry at each step and add it after the current node.
- If the sum have an extra carry at the end, simply make a new node and
  add it in the end.

Cost:
- O(n|m) time, O(m|n) space where m and m are lengths of these two lists.
*/

package leetcode

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestAddTwoNumbers(t *testing.T) {
	t11 := common.NewListNode(1)
	t11.AddNext(6)
	t11.AddNext(4)

	t12 := common.NewListNode(2)
	t12.AddNext(4)
	t12.AddNext(1)

	t13 := common.NewListNode(3)
	t13.AddNext(0)
	t13.AddNext(6)

	t21 := common.NewListNode(9)
	t21.AddNext(9)

	t22 := common.NewListNode(1)

	t23 := common.NewListNode(0)
	t23.AddNext(0)
	t23.AddNext(1)

	tests := []struct {
		in1      *common.ListNode
		in2      *common.ListNode
		expected *common.ListNode
	}{
		{t11, t12, t13},
		{t21, t22, t23},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			addTwoNumbers(tt.in1, tt.in2),
		)
	}
}

func addTwoNumbers(l1, l2 *common.ListNode) *common.ListNode {
	// initialize dummy heads to keep track of the head nodes.
	dummyHead := &common.ListNode{}
	current := dummyHead
	p1 := l1
	p2 := l2
	carry := 0

	// while either list is not empty.
	for p1 != nil || p2 != nil {
		// if p1 or p2 is nil, use 0 as its value.
		p1v, p2v := 0, 0
		if p1 != nil {
			p1v = p1.Value
		}
		if p2 != nil {
			p2v = p2.Value
		}

		// calculate the sum of corresponding digits and update it carry.
		sum := carry + p1v + p2v
		digit := sum % 10
		carry = sum / 10

		// add a next node the current node.
		current.Next = &common.ListNode{Value: digit, Next: nil}
		current = current.Next

		// keep traversing as long as we have not reached the end.
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}

	// if the sum have an extra carry at the end, simply make a new node and
	// add it in the end.
	if carry > 0 {
		current.Next = &common.ListNode{Value: carry, Next: nil}
	}

	return dummyHead.Next
}
