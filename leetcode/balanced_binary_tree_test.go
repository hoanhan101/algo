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
	// define test cases' input.
	t1 := &common.TreeNode{nil, 1, nil}

	t2 := &common.TreeNode{nil, 1, nil}
	t2.Right = &common.TreeNode{nil, 2, nil}

	t3 := &common.TreeNode{nil, 1, nil}
	t3.Right = &common.TreeNode{nil, 2, nil}
	t3.Right.Right = &common.TreeNode{nil, 3, nil}

	t4 := &common.TreeNode{nil, 1, nil}
	t4.Left = &common.TreeNode{nil, 2, nil}
	t4.Right = &common.TreeNode{nil, 3, nil}
	t4.Right.Right = &common.TreeNode{nil, 4, nil}

	t5 := &common.TreeNode{nil, 1, nil}
	t5.Left = &common.TreeNode{nil, 2, nil}
	t5.Right = &common.TreeNode{nil, 3, nil}
	t5.Right.Right = &common.TreeNode{nil, 4, nil}
	t5.Right.Right.Right = &common.TreeNode{nil, 5, nil}

	t6 := &common.TreeNode{nil, 1, nil}
	t6.Left = &common.TreeNode{nil, 2, nil}
	t6.Left.Left = &common.TreeNode{nil, 4, nil}
	t6.Left.Right = &common.TreeNode{nil, 5, nil}
	t6.Left.Right.Right = &common.TreeNode{nil, 6, nil}
	t6.Right = &common.TreeNode{nil, 3, nil}
	t6.Right.Right = &common.TreeNode{nil, 7, nil}
	t6.Right.Right.Left = &common.TreeNode{nil, 8, nil}
	t6.Right.Right.Right = &common.TreeNode{nil, 9, nil}
	t6.Right.Right.Right.Right = &common.TreeNode{nil, 10, nil}

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
	if maxDepth(t) != -1 {
		return true
	}

	return false
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
