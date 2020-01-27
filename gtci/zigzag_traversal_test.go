/*
Problem:
- Given a binary tree, populate the values of all nodes of each level
  in a zigzag order in separate sub-arrays.

Example:
- Input:
      1
	2   3
  4       5
  Output: [][]interface{}{[]interface{}{1}, []interface{}{3, 2}, []interface{}{4, 5}}

Approach:
- Start by pushing the root node to a queue.
- Keep iterating until the queue is empty.
- At each step,
  - use a linked list to push front or back depending on the zigzag direction
  - enqueue its left and right child

Cost:
- O(n) time, O(n) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestZigzagTraverse(t *testing.T) {
	t1 := &common.TreeNode{}

	t2 := &common.TreeNode{Left: nil, Value: 1, Right: nil}

	t3 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t3.Left = &common.TreeNode{Left: nil, Value: 2, Right: nil}
	t3.Left.Right = &common.TreeNode{Left: nil, Value: 3, Right: nil}

	t4 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t4.Left = &common.TreeNode{Left: nil, Value: 2, Right: nil}
	t4.Right = &common.TreeNode{Left: nil, Value: 3, Right: nil}

	t5 := &common.TreeNode{Left: nil, Value: 1, Right: nil}
	t5.Left = &common.TreeNode{Left: nil, Value: 2, Right: nil}
	t5.Right = &common.TreeNode{Left: nil, Value: 3, Right: nil}
	t5.Left.Left = &common.TreeNode{Left: nil, Value: 4, Right: nil}
	t5.Right.Right = &common.TreeNode{Left: nil, Value: 5, Right: nil}

	tests := []struct {
		in       *common.TreeNode
		expected [][]interface{}
	}{
		{t1, [][]interface{}{[]interface{}{0}}},
		{t2, [][]interface{}{[]interface{}{1}}},
		{t3, [][]interface{}{[]interface{}{1}, []interface{}{2}, []interface{}{3}}},
		{t4, [][]interface{}{[]interface{}{1}, []interface{}{3, 2}}},
		{t5, [][]interface{}{[]interface{}{1}, []interface{}{3, 2}, []interface{}{4, 5}}},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			zigzagTraverse(tt.in),
		)
	}
}

func zigzagTraverse(root *common.TreeNode) [][]interface{} {
	out := [][]interface{}{}

	if root == nil {
		return out
	}

	// initialize a linked list with the root.
	queue := common.NewQueue()
	queue.Push(root)
	leftToRight := true

	for queue.Size() > 0 {
		levelSize := queue.Size()
		currentLevel := common.NewList()

		for i := 0; i < levelSize; i++ {
			// pop the queue and cache that value to its current level.
			current := queue.Pop().(*common.TreeNode)

			if leftToRight {
				currentLevel.PushBack(current.Value)
			} else {
				currentLevel.PushFront(current.Value)
			}

			// push its left child.
			if current.Left != nil {
				queue.Push(current.Left)
			}

			// push its right child.
			if current.Right != nil {
				queue.Push(current.Right)
			}
		}

		out = append(out, currentLevel.Slice())

		// reverse its direction.
		leftToRight = !leftToRight
	}

	return out
}
