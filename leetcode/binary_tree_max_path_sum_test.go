/*
Problem:
- Given a binary tree, find the maximum path sum.

Assumption:
- The path might start and end at any node in the tree.
- Assume the tree is non-empty.
- The node can contain negative number.
- The maximum path does not have to go though the root node.

Approach:
- At each node, the potential maximum path could be one of these cases:
  - max(left subtree) + node
  - max(right subtree) + node
  - max(left subtree) + max(right subtree) + node
  - the node itself

Cost:
- O(n) time, O(n) space.
*/

package leetcode

import (
	"math"
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestGetMaxPathSum(t *testing.T) {
	t1 := &common.TreeNode{Left: nil, Value: 1, Right: nil}

	t2 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t2.Right = &common.TreeNode{Left: nil, Value: 2, Right: nil}

	t3 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t3.Left = &common.TreeNode{Left: nil, Value: 2, Right: nil}

	t4 := &common.TreeNode{Left: nil, Value: 5, Right: nil}
	t4.Left = &common.TreeNode{Left: nil, Value: 3, Right: nil}
	t4.Right = &common.TreeNode{Left: nil, Value: 8, Right: nil}

	t5 := &common.TreeNode{Left: nil, Value: 5, Right: nil}
	t5.Left = &common.TreeNode{Left: nil, Value: 3, Right: nil}
	t5.Right = &common.TreeNode{Left: nil, Value: 8, Right: nil}
	t5.Right.Left = &common.TreeNode{Left: nil, Value: 7, Right: nil}
	t5.Right.Right = &common.TreeNode{Left: nil, Value: 9, Right: nil}

	t6 := &common.TreeNode{Left: nil, Value: 5, Right: nil}
	t6.Left = &common.TreeNode{Left: nil, Value: 3, Right: nil}
	t6.Left.Left = &common.TreeNode{Left: nil, Value: 2, Right: nil}
	t6.Left.Left.Left = &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t6.Left.Right = &common.TreeNode{Left: nil, Value: 4, Right: nil}
	t6.Right = &common.TreeNode{Left: nil, Value: 8, Right: nil}
	t6.Right.Left = &common.TreeNode{Left: nil, Value: 7, Right: nil}
	t6.Right.Right = &common.TreeNode{Left: nil, Value: 9, Right: nil}
	t6.Right.Right.Right = &common.TreeNode{Left: nil, Value: 11, Right: nil}

	tests := []struct {
		in       *common.TreeNode
		expected int
	}{
		{t1, 1},
		{t2, 3},
		{t3, 3},
		{t4, 13},
		{t5, 22},
		{t6, 33},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			getMaxPathSum(tt.in),
		)
	}
}

// FIXME - need to look at this again to understand it and explain it better.
func getMaxPathSum(t *common.TreeNode) int {
	max := math.MinInt64

	if t == nil {
		return 0
	}

	leftMax := getMaxPathSum(t.Left)
	rightMax := getMaxPathSum(t.Right)
	max = common.Max(t.Value+leftMax+rightMax, max) //nolint
	m := common.Max(leftMax, rightMax)

	// s is the maximum path sum that goes through the current node and to
	// one if its left or right subtree to its parents.
	s := t.Value + m
	if s > 0 {
		return s
	}

	return 0
}
