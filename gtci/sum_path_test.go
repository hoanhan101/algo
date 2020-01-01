/*
Problem:
- Given a binary tree where each node can only have a digit (0-9) value,
  each root-to-leaf path will represent a number. Find the total sum of
  all the numbers represented by all paths.

Example:
- Input:
       1
    2     3
  4  5   6  7
  Output: 522 (= 124 + 125 + 136 + 137)

Approach:
- Traverse the tree in a depth first search fashion.
- At each level, the sum is equal to the result of the last sum times 10
  plus the current's node value.

Cost:
- O(n) time, O(n) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestSumPath(t *testing.T) {
	t1 := &common.TreeNode{nil, 1, nil}
	t1.Left = &common.TreeNode{nil, 2, nil}
	t1.Right = &common.TreeNode{nil, 3, nil}
	t1.Left.Left = &common.TreeNode{nil, 4, nil}
	t1.Left.Right = &common.TreeNode{nil, 5, nil}
	t1.Right.Left = &common.TreeNode{nil, 6, nil}
	t1.Right.Right = &common.TreeNode{nil, 7, nil}

	t2 := &common.TreeNode{nil, 1, nil}
	t2.Left = &common.TreeNode{nil, 2, nil}
	t2.Right = &common.TreeNode{nil, 3, nil}
	t2.Left.Left = &common.TreeNode{nil, 4, nil}
	t2.Right.Right = &common.TreeNode{nil, 7, nil}

	tests := []struct {
		in       *common.TreeNode
		expected int
	}{
		{t1, 522},
		{t2, 261},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			sumPath(tt.in),
		)
	}
}

func sumPath(root *common.TreeNode) int {
	return recurSumPath(root, 0)
}

func recurSumPath(node *common.TreeNode, sum int) int {
	if node == nil {
		return 0
	}

	sum = 10*sum + node.Value

	// return the current sum if the current node is the leaf node.
	if node.Left == nil && node.Right == nil {
		return sum
	}

	return recurSumPath(node.Left, sum) + recurSumPath(node.Right, sum)
}
