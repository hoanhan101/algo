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
	t1 := common.NewListNode(1)
	for i := 2; i <= 6; i++ {
		t1.AddNext(i)
	}

	// deletes node 4.
	deleteNode(t1.Next.Next.Next)
	common.Equal(t, []int{1, 2, 3, 5, 6}, common.LinkedListToSlice(t1))

	// deletes node 3.
	deleteNode(t1.Next.Next)
	common.Equal(t, []int{1, 2, 5, 6}, common.LinkedListToSlice(t1))

	// deletes node 5.
	deleteNode(t1.Next.Next)
	common.Equal(t, []int{1, 2, 6}, common.LinkedListToSlice(t1))
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
