/*
Problem:
- Given the head of a singly linked list, write a function to determine
  if it is a palindrome in constant space.

Approach:
- Find the middle of the linked list and reverse a half list
- After comparing the first half with the reversed half to check if it's
  a palindrome, revert to the half to original form.

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestIsPalindromeList(t *testing.T) {
	t1 := common.NewListNode(1)
	t1.AddNext(1)
	t1.AddNext(1)
	t1.AddNext(1)
	t1.AddNext(1)

	t2 := common.NewListNode(1)
	t2.AddNext(2)
	t2.AddNext(3)
	t2.AddNext(2)
	t2.AddNext(1)

	t3 := common.NewListNode(1)
	t3.AddNext(2)
	t3.AddNext(3)
	t3.AddNext(4)
	t3.AddNext(5)

	tests := []struct {
		in       *common.ListNode
		expected bool
	}{
		{t1, true},
		{t2, true},
		{t3, false},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			isPalindrome(tt.in),
		)
	}
}

func isPalindrome(head *common.ListNode) bool {
	slow, fast := head, head

	// find the middle of the linked list.
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

	}

	// reverse a half.
	headHalf := reverse(slow)
	copyHalf := headHalf

	// compare the first and second half.
	for head != nil && headHalf != nil {
		if head.Value != headHalf.Value {
			return false
		}

		head = head.Next
		headHalf = headHalf.Next
	}

	// revert the half to its original form.
	_ = reverse(copyHalf)

	if head == nil && headHalf == nil {
		return true
	}

	return false
}

func reverse(head *common.ListNode) *common.ListNode {
	var prev *common.ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
