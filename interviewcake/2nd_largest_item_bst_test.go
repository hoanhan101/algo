/*
Problem:
- Given a binary search tree, find the 2nd largest item.

Example:
- Input:
          5
      3       8
    2   4   7   9
  1               11
  Output: 9
- Input:
           5
      3        8
    2   4   7     16
                11
              9   12
  Output: 12

Approach:
- The largest item in a binary search tree is the rightmost item. Can
  simply traverse down the tree recursively to find one.
- The 2nd largest item could be the parent of the largest but it's not
  necessary since the largest could have a left subtree and there might exist
  one there.
- Still, the second largest one when the largest has a left subtree is basically
  the largest one in that left subtree.

Cost:
- O(h) time, O(1) space, where h is the height of the tree.
- If the tree is balanced, the time complexity is (Olgn). Otherwise, it's O(n).
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func Test2ndLargestItem(t *testing.T) {
	// define test cases' input.
	t1 := &BinaryTree{nil, 1, nil}

	t2 := &BinaryTree{nil, 1, nil}
	t2.right = &BinaryTree{nil, 2, nil}

	t3 := &BinaryTree{nil, 5, nil}
	t3.left = &BinaryTree{nil, 3, nil}
	t3.right = &BinaryTree{nil, 8, nil}

	t4 := &BinaryTree{nil, 5, nil}
	t4.left = &BinaryTree{nil, 3, nil}
	t4.right = &BinaryTree{nil, 8, nil}
	t4.right.left = &BinaryTree{nil, 7, nil}
	t4.right.right = &BinaryTree{nil, 9, nil}

	t5 := &BinaryTree{nil, 5, nil}
	t5.left = &BinaryTree{nil, 3, nil}
	t5.left.left = &BinaryTree{nil, 2, nil}
	t5.left.left.left = &BinaryTree{nil, 1, nil}
	t5.left.right = &BinaryTree{nil, 4, nil}
	t5.right = &BinaryTree{nil, 8, nil}
	t5.right.left = &BinaryTree{nil, 7, nil}
	t5.right.right = &BinaryTree{nil, 9, nil}
	t5.right.right.right = &BinaryTree{nil, 11, nil}

	t6 := &BinaryTree{nil, 5, nil}
	t6.left = &BinaryTree{nil, 3, nil}
	t6.right = &BinaryTree{nil, 8, nil}
	t6.right.left = &BinaryTree{nil, 7, nil}
	t6.right.right = &BinaryTree{nil, 16, nil}
	t6.right.right.left = &BinaryTree{nil, 11, nil}
	t6.right.right.left.left = &BinaryTree{nil, 9, nil}
	t6.right.right.left.right = &BinaryTree{nil, 12, nil}

	// define their outputs.
	tests := []struct {
		in       *BinaryTree
		expected int
	}{
		{t1, -1}, // doesn't have at least 2 nodes, return -1
		{t2, 1},
		{t3, 5},  // return the rightmost one
		{t4, 8},  // return the rightmost one
		{t5, 9},  // return the rightmost one
		{t6, 12}, // not the rightmost one, but the largest one in its subtree
	}

	for _, tt := range tests {
		result := findSecondLargest(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func findSecondLargest(t *BinaryTree) int {
	// if the tree is nil or doesn't have at least 2 nodes, return -1.
	if (t == nil) || (t.left == nil && t.right == nil) {
		return -1
	}

	current := t
	for current != nil {
		// if the current node is the largest node and it has a left subtree
		// then the 2nd largest one is the largest in that subtree.
		if (current.left != nil) && (current.right == nil) {
			return findLargest(current.left)
		}

		// if the current node is the parent of the largest node and the
		// largest has no children, it is the 2nd largest one.
		if (current.right != nil) && (current.right.right == nil) && (current.right.left == nil) {
			return current.value
		}

		current = current.right
	}

	return -1
}

// findLargest finds the largest item in a binary search tree.
func findLargest(t *BinaryTree) int {
	current := t

	for current != nil {
		// if there is no right child, we found the right most item so simply
		// return its value. otherwise, step down to this child and recurse.
		if current.right == nil {
			return current.value
		}
		current = current.right
	}

	return -1
}
