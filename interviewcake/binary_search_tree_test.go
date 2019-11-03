/*
Problem:
- Given a binary tree, determine if it is a binary search tree.

Example:
- Input:
          5
      3       8
    2   4   7   9
  1               11
  Output: true, because for each node, its value is greater than all values in
  the left subtree and less than all values in the right one.

Approach:
- Use a depth-first walk through the tree and validate each node as we go.
- If a node appears in the left subtree, it must be less than its ancestor and
  vice versa.
- Instead of keeping track of every ancestor to check the inequalities, just
  need to check the largest number it must be greater than and the smallest one
  it must be less than, aka lower bound and upper bound.

Solution:
- Initialize a stack that keeps track of the tree level, its lower bound and
  upper bound.
- Use an arbitrary low lower bound and high higher bound to start the stack.
- While the stack is not empty, pop a tree and its bounds from the top.
- Return false immediately if the current node is invalid, meaning its value
  is either:
  - less or equal than the lower bound
  - greater or equal than upper bound
- Keep walking down the tree and update the lower and upper bounds accordingly.

Cost:
- O(n) time, O(n) space.
- The worst case is that we have to iterate all nodes in the tree so the time
  complexity is O(n). For space complexity, we have to keep track of the lower
  bound and upper bound as we traverse the tree via a stack. Hence, the worst
  case is O(n).
*/

package interviewcake

import (
	"math"
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestIsBinarySearchTree(t *testing.T) {
	// define test cases' input.
	t1 := &BinaryTree{nil, 1, nil}

	t2 := &BinaryTree{nil, 1, nil}
	t2.right = &BinaryTree{nil, 2, nil}

	t3 := &BinaryTree{nil, 1, nil}
	t3.left = &BinaryTree{nil, 2, nil}

	t4 := &BinaryTree{nil, 5, nil}
	t4.left = &BinaryTree{nil, 3, nil}
	t4.right = &BinaryTree{nil, 8, nil}

	t5 := &BinaryTree{nil, 5, nil}
	t5.left = &BinaryTree{nil, 3, nil}
	t5.right = &BinaryTree{nil, 8, nil}
	t5.right.left = &BinaryTree{nil, 7, nil}
	t5.right.right = &BinaryTree{nil, 9, nil}

	t6 := &BinaryTree{nil, 5, nil}
	t6.left = &BinaryTree{nil, 3, nil}
	t6.left.left = &BinaryTree{nil, 2, nil}
	t6.left.left.left = &BinaryTree{nil, 1, nil}
	t6.left.right = &BinaryTree{nil, 4, nil}
	t6.right = &BinaryTree{nil, 8, nil}
	t6.right.left = &BinaryTree{nil, 7, nil}
	t6.right.right = &BinaryTree{nil, 9, nil}
	t6.right.right.right = &BinaryTree{nil, 11, nil}

	// define their outputs.
	tests := []struct {
		in       *BinaryTree
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

// treeBound holds the tree, its lower bound and upper bound.
type treeBound struct {
	tree       *BinaryTree
	lowerBound int
	upperBound int
}

func isBinarySearchTree(t *BinaryTree) bool {
	// stack keeps track of the tree level, its lower bound and upper bound.
	stack := []treeBound{}

	// start at the root with arbitrary low lower bound and high higher bound.
	stack = append(stack, treeBound{t, math.MinInt8, math.MaxInt8})

	for len(stack) > 0 {
		// pop a tree and its bounds from the top of our stack.
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// return false immediately if the currnent node is invalid, meaning
		// its value is either:
		// - less or equal than the lower bound
		// - greater or equal than upper bound
		if current.tree.value <= current.lowerBound || current.tree.value >= current.upperBound {
			return false
		}

		// keep walking down the tree and keep track of the lower and upper
		// bounds.
		if current.tree.left != nil {
			// since this node must be less than the current node value, let
			// upper bound be the current node value.
			stack = append(stack, treeBound{current.tree.left, current.lowerBound, current.tree.value})
		}
		if current.tree.right != nil {
			// since this node must be greater than the current node value, let
			// lower bound be the current node value.
			stack = append(stack, treeBound{current.tree.right, current.tree.value, current.upperBound})
		}
	}

	// at this point we've checked all nodes and since they are all valid, it
	// is a binary search tree.
	return true
}
