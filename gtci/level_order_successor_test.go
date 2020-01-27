/*
Problem:
- Given a binary tree and a node, find the level order successor of
  the given node. The level order successor is the node that appears
  right after the given node in the level order traversal.

Example:
- Input: target=6, tree=
      1
	2   3
  4       5
        6   7
      8
  Output: 6

Approach:
- Similar to level order traversal problem, except we will not keep
  track of all the levels.
- Instead, we keep inserting the child node to the queue and return
  the next node as soon as we find the target.

Cost:
- O(n) time, O(n) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFindSuccessor(t *testing.T) {
	t1 := &common.TreeNode{Left: nil, Value: 1, Right: nil}

	t2 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t2.Left = &common.TreeNode{Left: nil, Value: 2, Right: nil}
	t2.Left.Right = &common.TreeNode{Left: nil, Value: 3, Right: nil}

	t3 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t3.Left = &common.TreeNode{Left: nil, Value: 2, Right: nil}
	t3.Right = &common.TreeNode{Left: nil, Value: 3, Right: nil}

	t4 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t4.Left = &common.TreeNode{Left: nil, Value: 2, Right: nil}
	t4.Right = &common.TreeNode{Left: nil, Value: 3, Right: nil}
	t4.Left.Left = &common.TreeNode{Left: nil, Value: 4, Right: nil}
	t4.Right.Right = &common.TreeNode{Left: nil, Value: 5, Right: nil}

	t5 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t5.Left = &common.TreeNode{Left: nil, Value: 2, Right: nil}
	t5.Right = &common.TreeNode{Left: nil, Value: 3, Right: nil}
	t5.Left.Left = &common.TreeNode{Left: nil, Value: 4, Right: nil}
	t5.Right.Left = &common.TreeNode{Left: nil, Value: 5, Right: nil}
	t5.Right.Right = &common.TreeNode{Left: nil, Value: 6, Right: nil}
	t5.Left.Left.Left = &common.TreeNode{Left: nil, Value: 7, Right: nil}
	t5.Left.Left.Left.Left = &common.TreeNode{Left: nil, Value: 8, Right: nil}

	tests := []struct {
		in1      *common.TreeNode
		in2      int
		expected int
	}{
		{t1, 1, 0},
		{t2, 2, 3},
		{t3, 2, 3},
		{t4, 4, 5},
		{t5, 6, 7},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			findSuccessor(tt.in1, tt.in2),
		)
	}
}

func findSuccessor(root *common.TreeNode, target int) int {
	if root == nil {
		return 0
	}

	// initialize a linked list with the root.
	queue := common.NewQueue()
	queue.Push(root)

	for queue.Size() > 0 {
		// pop the queue and cache the current value.
		current := queue.Pop().(*common.TreeNode)

		// push its left child.
		if current.Left != nil {
			queue.Push(current.Left)
		}

		// push its right child.
		if current.Right != nil {
			queue.Push(current.Right)
		}

		if current.Value == target {
			break
		}

	}

	if queue.Size() > 0 {
		return queue.Pop().(*common.TreeNode).Value
	}

	return 0
}
