/*
Problem:
- Given a binary tree, find its maximum depth.

Approach:
- The maximum depth of the current node is the greater of the max height of the left
  subtree and the right subtree plus one.

Cost:
- O(n) time, O(n) space.
*/

package leetcode

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestGetMaxDepth(t *testing.T) {
	// define test cases' input.
	t1 := &common.TreeNode{nil, 1, nil}

	t2 := &common.TreeNode{nil, 1, nil}
	t2.Right = &common.TreeNode{nil, 2, nil}

	t3 := &common.TreeNode{nil, 1, nil}
	t3.Left = &common.TreeNode{nil, 2, nil}

	t4 := &common.TreeNode{nil, 5, nil}
	t4.Left = &common.TreeNode{nil, 3, nil}
	t4.Right = &common.TreeNode{nil, 8, nil}

	t5 := &common.TreeNode{nil, 5, nil}
	t5.Left = &common.TreeNode{nil, 3, nil}
	t5.Right = &common.TreeNode{nil, 8, nil}
	t5.Right.Left = &common.TreeNode{nil, 7, nil}
	t5.Right.Right = &common.TreeNode{nil, 9, nil}

	t6 := &common.TreeNode{nil, 5, nil}
	t6.Left = &common.TreeNode{nil, 3, nil}
	t6.Left.Left = &common.TreeNode{nil, 2, nil}
	t6.Left.Left.Left = &common.TreeNode{nil, 1, nil}
	t6.Left.Right = &common.TreeNode{nil, 4, nil}
	t6.Right = &common.TreeNode{nil, 8, nil}
	t6.Right.Left = &common.TreeNode{nil, 7, nil}
	t6.Right.Right = &common.TreeNode{nil, 9, nil}
	t6.Right.Right.Right = &common.TreeNode{nil, 11, nil}

	tests := []struct {
		in       *common.TreeNode
		expected int
	}{
		{t1, 1},
		{t2, 2},
		{t3, 2},
		{t4, 2},
		{t5, 3},
		{t6, 4},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			getMaxDepth(tt.in),
		)
	}
}

func getMaxDepth(t *common.TreeNode) int {
	// for the base case, the height is 0.
	if t == nil {
		return 0
	}

	// calculate the max depth of left subtree and right subtree.
	maxLeft := getMaxDepth(t.Left)
	maxRight := getMaxDepth(t.Right)

	// the max depth of the tree is the greater one plus one.
	_, max := common.Mimax(maxLeft, maxRight)
	return max + 1
}
