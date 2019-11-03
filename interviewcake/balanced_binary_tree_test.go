/*
Problem:
- Given a binary tree, determine if it is "superbalanced" - the difference
  between the depths of any two leaf nodes is no greater than 1.

Example:
- Input:
         1
      2     3
    4   5      7
          6  8   9
                   10
  Output: false
  Even though this tree is balanced by definition, it is not "superbalanced".

Approach:
- Use a depth-first walk through the tree and keep track of the depth as we
  go.
- Every time we found a leaf with a new depth, there are two ways that the
  tree could be unbalanced:
   - There are more than 2 different leaf depths.
   - There are exactly 2 depths but they are more than 1 apart.

Solution:
- Initialize a stack to keep track of the tree level and its depth.
- While the stack is not empty, pop a tree and its depth from the top
  of our stack and check if we found a leaf.
- If so, add the new depth to the list if we haven't seen it.
- Could short-circuit to determine if the tree is unbalanced if:
  - more than 2 different leaf depths
  - 2 leaf depths that are more than 1 apart
- Keep walking down the tree and keep track of the depth.

Cost:
- O(n) time, O(n) space.
- The worst case is that we have to iterate all nodes in the tree so the time
  complexity is O(n). For space complexity, we have to keep track of the all
  the nodes at every depth. Hence, it is O(n).
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestIsSuperBalanced(t *testing.T) {
	// define test cases' input.
	t1 := &BinaryTree{nil, 1, nil}

	t2 := &BinaryTree{nil, 1, nil}
	t2.right = &BinaryTree{nil, 2, nil}

	t3 := &BinaryTree{nil, 1, nil}
	t3.right = &BinaryTree{nil, 2, nil}
	t3.right.right = &BinaryTree{nil, 3, nil}

	t4 := &BinaryTree{nil, 1, nil}
	t4.left = &BinaryTree{nil, 2, nil}
	t4.right = &BinaryTree{nil, 3, nil}
	t4.right.right = &BinaryTree{nil, 4, nil}

	t5 := &BinaryTree{nil, 1, nil}
	t5.left = &BinaryTree{nil, 2, nil}
	t5.right = &BinaryTree{nil, 3, nil}
	t5.right.right = &BinaryTree{nil, 4, nil}
	t5.right.right.right = &BinaryTree{nil, 5, nil}

	t6 := &BinaryTree{nil, 1, nil}
	t6.left = &BinaryTree{nil, 2, nil}
	t6.left.left = &BinaryTree{nil, 4, nil}
	t6.left.right = &BinaryTree{nil, 5, nil}
	t6.left.right.right = &BinaryTree{nil, 6, nil}
	t6.right = &BinaryTree{nil, 3, nil}
	t6.right.right = &BinaryTree{nil, 7, nil}
	t6.right.right.left = &BinaryTree{nil, 8, nil}
	t6.right.right.right = &BinaryTree{nil, 9, nil}
	t6.right.right.right.right = &BinaryTree{nil, 10, nil}

	// define their outputs.
	tests := []struct {
		in       *BinaryTree
		expected bool
	}{
		{t1, true},
		{t2, true},
		{t3, true},
		{t4, true},
		{t5, false},
		{t6, false},
	}

	for _, tt := range tests {
		result := isSuperBalanced(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

// BinaryTree represents a binary tree.
type BinaryTree struct {
	left  *BinaryTree
	value int
	right *BinaryTree
}

// treeDepth holds the tree and its depth level.
type treeDepth struct {
	tree  *BinaryTree
	depth int
}

func isSuperBalanced(t *BinaryTree) bool {
	// return true if the tree has no leaf.
	if t == nil {
		return true
	}

	// depths holds a list of depth that we have seen.
	depths := []int{}

	// stack keeps track of the tree level and its depth.
	stack := []treeDepth{}
	stack = append(stack, treeDepth{t, 0})

	for len(stack) > 0 {
		// pop a tree level and its depth from the top of our stack.
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// if we found a leaf, add the new depth to the list if we haven't seen it.
		if current.tree.left == nil && current.tree.right == nil {
			if !common.Contain(depths, current.depth) {
				depths = append(depths, current.depth)
			}

			// short-circuit to determine if the tree is unbalanced:
			// - more than 2 different leaf depths
			// - 2 leaf depths that are more than 1 apart
			if (len(depths) > 2) || (len(depths) == 2 && common.IsMoreThan1Apart(depths[1], depths[0])) {
				return false
			}
		}

		// keep walking down the tree and keep track of the depth.
		if current.tree.left != nil {
			stack = append(stack, treeDepth{current.tree.left, current.depth + 1})
		}
		if current.tree.right != nil {
			stack = append(stack, treeDepth{current.tree.right, current.depth + 1})
		}
	}

	// at this point we've checked all nodes and since they are all valid, the
	// tree is "superbalanced".
	return true
}
