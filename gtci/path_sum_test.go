/*
Problem:
- Given a binary tree and a sum, find if the tree has a path from
  root-to-leaf such that the sum of all the node values of that path
  equals to the sum.

Example:
- Input: sum=8
       1
    2     3
  4  5   6  7
  Output: true
  Explanation: The path with sum 8 goes through 1, 2 and 5.
- Input: sum=16
       1
    2     3
  4  5   6  7
  Output: false
  Explanation: There is no path with sum 16.

Approach:
- Traverse the tree in a depth first search fashion.
- At each step, update the new sum by subtracting it with the current
  value of the node's we're visiting.
- Once we hit the leaf node, return true if the sum is equal to its
  value.

Cost:
- O(n) time, O(n) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestPathSum(t *testing.T) {
	t1 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t1.Left = &common.TreeNode{Left: nil, Value: 2, Right: nil}
	t1.Right = &common.TreeNode{Left: nil, Value: 3, Right: nil}
	t1.Left.Left = &common.TreeNode{Left: nil, Value: 4, Right: nil}
	t1.Left.Right = &common.TreeNode{Left: nil, Value: 5, Right: nil}
	t1.Right.Left = &common.TreeNode{Left: nil, Value: 6, Right: nil}
	t1.Right.Right = &common.TreeNode{Left: nil, Value: 7, Right: nil}

	tests := []struct {
		in1      *common.TreeNode
		in2      int
		expected bool
	}{
		{t1, 7, true},
		{t1, 8, true},
		{t1, 10, true},
		{t1, 11, true},
		{t1, 16, false},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			pathSum(tt.in1, tt.in2),
		)
	}
}

func pathSum(root *common.TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Value == sum && root.Left == nil && root.Right == nil {
		return true
	}

	return pathSum(root.Left, sum-root.Value) || pathSum(root.Right, sum-root.Value)
}
