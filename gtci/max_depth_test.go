/*
Problem:
- Given a binary tree, find the maximum depth.

Example:
- Input:
      1
	2   3
  4       5
        6   7
      8
  Output: 5

Approach:
- Similar to minimum depth problem, we will keep traversing for all
  levels, incrementing our maximum depth instead of returning as
  soon as we find a leaf node.

Cost:
- O(n) time, O(n) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestMaxDepth(t *testing.T) {
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
		in       *common.TreeNode
		expected int
	}{
		{t1, 1},
		{t2, 3},
		{t3, 2},
		{t4, 3},
		{t5, 5},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			maxDepth(tt.in),
		)
	}
}

func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	// initialize a linked list with the root.
	queue := common.NewQueue()
	queue.Push(root)

	// track the minimum depth.
	maxDepth := 0

	for queue.Size() > 0 {
		// increase the max depth at each level.
		maxDepth++

		levelSize := queue.Size()

		for i := 0; i < levelSize; i++ {
			// pop the queue and cache that value to its current level.
			current := queue.Pop().(*common.TreeNode)

			// push its left child.
			if current.Left != nil {
				queue.Push(current.Left)
			}

			// push its right child.
			if current.Right != nil {
				queue.Push(current.Right)
			}
		}

	}

	return maxDepth
}
