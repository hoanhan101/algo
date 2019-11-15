/*
Problem:
- Delete a node from a singly-linked list, given only a pointer to that node.

Approach:
- Since we don't have access to the previous node, simply copy the value and
  pointer of the next node and copy them into the current node.

Solution:
- Cache the next node.
- If the next node is nil, it's the last node. Just simply return.
- Copy the current node's value to the next node's value
- Copy the node's pointer to the next node's pointer.

Cost:
- O(1) time and O(1) space.
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestDeleteNode(t *testing.T) {
	// define test input.
	t1 := &common.ListNode{1, nil}
	t1.Next = &common.ListNode{2, nil}
	t1.Next.Next = &common.ListNode{3, nil}
	t1.Next.Next.Next = &common.ListNode{4, nil}
	t1.Next.Next.Next.Next = &common.ListNode{5, nil}
	t1.Next.Next.Next.Next.Next = &common.ListNode{6, nil}

	// deletes node 4.
	deleteNode(t1.Next.Next.Next)
	common.Equal(t, []int{1, 2, 3, 5, 6}, linkedListToSlice(t1))

	// deletes node 3.
	deleteNode(t1.Next.Next)
	common.Equal(t, []int{1, 2, 5, 6}, linkedListToSlice(t1))

	// deletes node 5.
	deleteNode(t1.Next.Next)
	common.Equal(t, []int{1, 2, 6}, linkedListToSlice(t1))
}

func deleteNode(node *common.ListNode) {
	nextNode := node.Next

	// if the next node is nil, it's the last node. this implementation does
	// not work.
	if nextNode == nil {
		return
	}

	node.Value = nextNode.Value
	node.Next = nextNode.Next
}

// linkedListToSlice converts a linked list into an array of integer.
func linkedListToSlice(node *common.ListNode) []int {
	out := []int{}

	for n := node; n != nil; n = n.Next {
		out = append(out, n.Value)
	}

	return out
}
