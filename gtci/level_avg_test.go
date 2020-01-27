/*
Problem:
- Given a binary tree, populate an array to represent the averages of all of
  its levels.

Example:
- Input:
      1
	2   3
  4       5
  Output: []float64{1, 2.5, 4.5}

Approach:
- Similar to level order traversal problem, except we keep track of the sum
  at each level and return the average in the end.

Cost:
- O(n) time, O(n) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestLevelAvg(t *testing.T) {
	t1 := &common.TreeNode{}

	t2 := &common.TreeNode{Left: nil, Value: 1, Right: nil}

	t3 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t3.Left = &common.TreeNode{Left: nil, Value: 2, Right: nil}
	t3.Left.Right = &common.TreeNode{Left: nil, Value: 3, Right: nil}

	t4 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t4.Left = &common.TreeNode{Left: nil, Value: 2, Right: nil}
	t4.Right = &common.TreeNode{Left: nil, Value: 3, Right: nil}

	t5 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t5.Left = &common.TreeNode{Left: nil, Value: 2, Right: nil}
	t5.Right = &common.TreeNode{Left: nil, Value: 3, Right: nil}
	t5.Left.Left = &common.TreeNode{Left: nil, Value: 4, Right: nil}
	t5.Right.Right = &common.TreeNode{Left: nil, Value: 5, Right: nil}

	tests := []struct {
		in       *common.TreeNode
		expected []float64
	}{
		{t1, []float64{0}},
		{t2, []float64{1}},
		{t3, []float64{1, 2, 3}},
		{t4, []float64{1, 2.5}},
		{t5, []float64{1, 2.5, 4.5}},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			levelAvg(tt.in),
		)
	}
}

func levelAvg(root *common.TreeNode) []float64 {
	out := []float64{}

	if root == nil {
		return out
	}

	// initialize a linked list with the root.
	queue := common.NewQueue()
	queue.Push(root)

	for queue.Size() > 0 {
		levelSize := queue.Size()
		levelSum := 0

		for i := 0; i < levelSize; i++ {
			// pop the queue and cache that value to its current level.
			current := queue.Pop().(*common.TreeNode)

			levelSum += current.Value

			// push its left child.
			if current.Left != nil {
				queue.Push(current.Left)
			}

			// push its right child.
			if current.Right != nil {
				queue.Push(current.Right)
			}
		}

		out = append(out, float64(levelSum)/float64(levelSize))
	}

	return out
}
