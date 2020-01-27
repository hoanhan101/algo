/*
Problem:
- Given a binary tree, determine if it is a valid binary search tree.

Approach:
- Traverse the tree and apply recursion to check at each step if:
  - the current node's value is greater than the lower bound
  - the current node's value is smaller than the upper bound
  - the current node's left child follows
  - the current node's left child follows

Cost:
- O(n) time and O(n) stack space.
*/

package leetcode

import (
	"math"
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestIsBinarySearchTree(t *testing.T) {
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

	// define their outputs.
	tests := []struct {
		in       *common.TreeNode
		expected bool
	}{
		{t1, true},
		{t2, true},
		{t3, false},
		{t4, true},
		{t5, true},
		{t6, true},
	}

	for _, tt := range tests {
		result := isBinarySearchTree(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func isBinarySearchTree(t *common.TreeNode) bool {
	return validate(t, math.MinInt64, math.MaxInt64)
}

func validate(t *common.TreeNode, lower, upper int) bool {
	if t == nil {
		return true
	}

	return t.Value > lower && t.Value < upper && validate(t.Left, lower, t.Value) && validate(t.Right, t.Value, upper)
}
