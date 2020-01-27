/*
Problem:
- Given a binary tree, find its minimum depth.

Approach:
- Similar to finding maximum depth, the minimum depth of the current node is
  the smaller of the min height of the left subtree and the right subtree plus one.

Cost:
- O(n) time, O(n) space.
*/

package leetcode

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestGetMinDepth(t *testing.T) {
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
		{t2, 2},
		{t3, 2},
		{t4, 2},
		{t5, 2},
		{t6, 3},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			getMinDepth(tt.in),
		)
	}
}

func getMinDepth(t *common.TreeNode) int {
	// for the base case, the height is 0.
	if t == nil {
		return 0
	}

	// if a node has an empty subtree, return the minimum depth of the other
	// non-empty subtree.
	if t.Left == nil {
		return getMinDepth(t.Right) + 1
	}

	if t.Right == nil {
		return getMinDepth(t.Left) + 1
	}

	// calculate the min depth of left subtree and right subtree.
	minLeft := getMinDepth(t.Left)
	minRight := getMinDepth(t.Right)

	// the max depth of the tree is the greater one plus one.
	return common.Min(minLeft, minRight) + 1
}
