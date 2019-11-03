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
	t1 := &LinkedList{1, nil}
	t1.next = &LinkedList{2, nil}
	t1.next.next = &LinkedList{3, nil}
	t1.next.next.next = &LinkedList{4, nil}
	t1.next.next.next.next = &LinkedList{5, nil}
	t1.next.next.next.next.next = &LinkedList{6, nil}

	// deletes node 4.
	deleteNode(t1.next.next.next)
	common.Equal(t, []int{1, 2, 3, 5, 6}, linkedListToSlice(t1))

	// deletes node 3.
	deleteNode(t1.next.next)
	common.Equal(t, []int{1, 2, 5, 6}, linkedListToSlice(t1))

	// deletes node 5.
	deleteNode(t1.next.next)
	common.Equal(t, []int{1, 2, 6}, linkedListToSlice(t1))
}

// LinkedList represents a node in a linked list.
type LinkedList struct {
	value int
	next  *LinkedList
}

func deleteNode(node *LinkedList) {
	nextNode := node.next

	// if the next node is nil, it's the last node. this implementation does
	// not work.
	if nextNode == nil {
		return
	}

	node.value = nextNode.value
	node.next = nextNode.next
}

// linkedListToSlice converts a linked list into an array of integer.
func linkedListToSlice(node *LinkedList) []int {
	out := []int{}

	for n := node; n != nil; n = n.next {
		out = append(out, n.value)
	}

	return out
}
