/*
Problem:
- Determine if a binary tree is height-balanced.

Example:
- Input:
         1
      2     3
    4
  Output: true

- Input:
         1
      2     3
    4
      5
  Output: false

Approach:
A binary is balanced if
- its left subtree if balanced
- its right subtree if balanced
- the difference between the heights of left subtree and right subtree is
not more than 1.
*/

package other

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestIsBalanced(t *testing.T) {
	t1 := &BinaryTree{nil, 1, nil}

	t2 := &BinaryTree{nil, 1, nil}
	t2.left = &BinaryTree{nil, 2, nil}

	t3 := &BinaryTree{nil, 1, nil}
	t3.left = &BinaryTree{nil, 2, nil}
	t3.right = &BinaryTree{nil, 3, nil}
	t3.right.right = &BinaryTree{nil, 4, nil}

	t4 := &BinaryTree{nil, 1, nil}
	t4.right = &BinaryTree{nil, 2, nil}
	t4.right.right = &BinaryTree{nil, 3, nil}
	t4.right.right.right = &BinaryTree{nil, 4, nil}

	t5 := &BinaryTree{nil, 1, nil}
	t5.left = &BinaryTree{nil, 2, nil}
	t5.right = &BinaryTree{nil, 3, nil}
	t5.left.left = &BinaryTree{nil, 4, nil}

	t6 := &BinaryTree{nil, 1, nil}
	t6.left = &BinaryTree{nil, 2, nil}
	t6.right = &BinaryTree{nil, 3, nil}
	t6.left.left = &BinaryTree{nil, 4, nil}
	t6.left.left.right = &BinaryTree{nil, 5, nil}

	tests := []struct {
		in       *BinaryTree
		height   int
		balanced bool
	}{
		{t1, 1, true},
		{t2, 2, true},
		{t3, 3, true},
		{t4, 4, false},
		{t5, 3, true},
		{t6, 4, false},
	}

	for _, tt := range tests {
		h := height(tt.in)
		b := isBalanced(tt.in)
		common.Equal(t, tt.height, h)
		common.Equal(t, tt.balanced, b)
	}
}

func isBalanced(t *BinaryTree) bool {
	if t == nil {
		return true
	}

	// calculate the left and right subtrees' heights.
	leftHeight := height(t.left)
	rightHeight := height(t.right)

	// if the difference between left and right subtrees' heights is less than
	// 1 apart while the left and right subtree are balanced then the tree must
	// be balanced.
	if common.IsLessThan1Apart(leftHeight, rightHeight) && isBalanced(t.left) && isBalanced(t.right) {
		return true
	}

	return false
}

// height returns the height of the binary tree.
func height(t *BinaryTree) int {
	if t == nil {
		return 0
	}

	return common.Max(height(t.left), height(t.right)) + 1
}
