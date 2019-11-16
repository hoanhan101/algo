/*
Problem:
- Merge two sorted linked lists and return it as a new list.

Example:
- Input: 1 -> 3-> 5 & 2 -> 4-> 6
  Output: 1 -> 2-> 3 -> 4 -> 5 -> 6

Approach:
- Traverse both list at the same time, compare their values at each step and
  add the smaller one to a new list.

Solution:
- Initialize a dummy head to keep track of the head node.
- Traverse both lists while they are are not empty.
- Compare their values at each step and add the smaller one to a new list.
- Remember to check if we have reached the end for a list faster than the
  other.
- If that is the case, simply add the rest of the other list to the new list
  since it is already sorted.

Cost:
- O(n|m) time, O(n+m) space where n and m are lengths of these two linked lists.
*/

package leetcode

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestMergeSortedLinkedList(t *testing.T) {
	t11 := common.NewListNode(1)
	t11.AddNext(3)
	t11.AddNext(5)

	t12 := common.NewListNode(2)
	t12.AddNext(4)
	t12.AddNext(6)

	t13 := common.NewListNode(1)
	for i := 2; i <= 6; i++ {
		t13.AddNext(i)
	}

	tests := []struct {
		in1      *common.ListNode
		in2      *common.ListNode
		expected *common.ListNode
	}{
		{t11, t12, t13},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			mergeSortedLinkedList(tt.in1, tt.in2),
		)
	}
}

func mergeSortedLinkedList(l1, l2 *common.ListNode) *common.ListNode {
	// initialize a dummy head to keep track of the head node.
	dummyHead := &common.ListNode{}
	current := dummyHead

	// while both linked list are not empty.
	for l1 != nil && l2 != nil {
		if l1.Value < l2.Value {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}

		current = current.Next
	}

	// if we are here, it means that we have reached the end of a linked list.
	// simply append the rest of the other list since it's already sorted.
	if l1 != nil {
		current.Next = l1
	}

	if l2 != nil {
		current.Next = l2
	}

	return dummyHead.Next
}
