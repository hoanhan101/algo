/*
Problem:
- Given a binary tree, populate the values of all nodes of each level
  from left to right in separate sub-arrays.

Example:
- Input:
      1
	2   3
  4       5
  Output: [][]int{[]int{1}, []int{2, 3}, []int{4, 5}}

Approach:
- Start by pushing the root node to the queue.
- Keep iterating until the queue is empty.
- At each step,
  - send the front of the queue and dequeue it
  - enqueue its left and right child

Cost:
- O(n) time, O(n) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestTraverse(t *testing.T) {
	t1 := &common.TreeNode{}

	t2 := &common.TreeNode{nil, 1, nil}

	t3 := &common.TreeNode{nil, 1, nil}
	t3.Left = &common.TreeNode{nil, 2, nil}
	t3.Left.Right = &common.TreeNode{nil, 3, nil}

	t4 := &common.TreeNode{nil, 1, nil}
	t4.Left = &common.TreeNode{nil, 2, nil}
	t4.Right = &common.TreeNode{nil, 3, nil}

	t5 := &common.TreeNode{nil, 1, nil}
	t5.Left = &common.TreeNode{nil, 2, nil}
	t5.Right = &common.TreeNode{nil, 3, nil}
	t5.Left.Left = &common.TreeNode{nil, 4, nil}
	t5.Right.Right = &common.TreeNode{nil, 5, nil}

	tests := []struct {
		in       *common.TreeNode
		expected [][]int
	}{
		{t1, [][]int{[]int{0}}},
		{t2, [][]int{[]int{1}}},
		{t3, [][]int{[]int{1}, []int{2}, []int{3}}},
		{t4, [][]int{[]int{1}, []int{2, 3}}},
		{t5, [][]int{[]int{1}, []int{2, 3}, []int{4, 5}}},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			traverse(tt.in),
		)
	}
}

func traverse(root *common.TreeNode) [][]int {
	out := [][]int{}

	if root == nil {
		return out
	}

	// initialize a queue with the root.
	queue := common.NewQueue()
	queue.Push(root)

	for queue.Size() > 0 {
		levelSize := queue.Size()
		currentLevel := []int{}

		for i := 0; i < levelSize; i++ {
			// pop the queue and cache that value to its current level.
			current := queue.Pop().(*common.TreeNode)
			currentLevel = append(currentLevel, current.Value)

			// push its left child.
			if current.Left != nil {
				queue.Push(current.Left)
			}

			// push its right child.
			if current.Right != nil {
				queue.Push(current.Right)
			}
		}

		out = append(out, currentLevel)
	}

	return out
}
