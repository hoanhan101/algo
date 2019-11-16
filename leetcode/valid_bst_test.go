/*
Problem:
- Given a binary tree, determine if it is a valid binary search tree (BST).

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
