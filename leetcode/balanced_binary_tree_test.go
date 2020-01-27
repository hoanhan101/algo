/*
Problem:
- Given a binary tree, determine if it is height-balanced.

Approach:
- Calculate max depth for the left subtree and right subtree.
- If either the left subtree or right subtree is unbalanced, return right away.

Cost:
- O(n) time, O(n) stack space.
*/

package leetcode

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestIsBalanced(t *testing.T) {
	t1 := &common.TreeNode{Left: nil, Value: 1, Right: nil}

	t2 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t2.Right = &common.TreeNode{Left: nil, Value: 2, Right: nil}

	t3 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t3.Right = &common.TreeNode{Left: nil, Value: 2, Right: nil}
	t3.Right.Right = &common.TreeNode{Left: nil, Value: 3, Right: nil}

	t4 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t4.Left = &common.TreeNode{Left: nil, Value: 2, Right: nil}
	t4.Right = &common.TreeNode{Left: nil, Value: 3, Right: nil}
	t4.Right.Right = &common.TreeNode{Left: nil, Value: 4, Right: nil}

	t5 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t5.Left = &common.TreeNode{Left: nil, Value: 2, Right: nil}
	t5.Right = &common.TreeNode{Left: nil, Value: 3, Right: nil}
	t5.Right.Right = &common.TreeNode{Left: nil, Value: 4, Right: nil}
	t5.Right.Right.Right = &common.TreeNode{Left: nil, Value: 5, Right: nil}

	t6 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t6.Left = &common.TreeNode{Left: nil, Value: 2, Right: nil}
	t6.Left.Left = &common.TreeNode{Left: nil, Value: 4, Right: nil}
	t6.Left.Right = &common.TreeNode{Left: nil, Value: 5, Right: nil}
	t6.Left.Right.Right = &common.TreeNode{Left: nil, Value: 6, Right: nil}
	t6.Right = &common.TreeNode{Left: nil, Value: 3, Right: nil}
	t6.Right.Right = &common.TreeNode{Left: nil, Value: 7, Right: nil}
	t6.Right.Right.Left = &common.TreeNode{Left: nil, Value: 8, Right: nil}
	t6.Right.Right.Right = &common.TreeNode{Left: nil, Value: 9, Right: nil}
	t6.Right.Right.Right.Right = &common.TreeNode{Left: nil, Value: 10, Right: nil}

	tests := []struct {
		in       *common.TreeNode
		expected bool
	}{
		{t1, true},
		{t2, true},
		{t3, false},
		{t4, true},
		{t5, false},
		{t6, false},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			isBalanced(tt.in),
		)
	}
}

func isBalanced(t *common.TreeNode) bool {
	return maxDepth(t) != -1
}

func maxDepth(t *common.TreeNode) int {
	if t == nil {
		return 0
	}

	maxLeft := maxDepth(t.Left)
	if maxLeft == -1 {
		return -1
	}

	maxRight := maxDepth(t.Right)
	if maxRight == -1 {
		return -1
	}

	if common.IsLessThan1Apart(maxLeft, maxRight) {
		return common.Max(maxLeft, maxRight) + 1
	}

	return -1
}
