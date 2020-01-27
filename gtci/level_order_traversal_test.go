/*
Problem:
- Given a binary tree, populate the values of all nodes of each level
  from left to right in separate sub-arrays.

Example:
- Input:
      1
	2   3
  4       5
  Output: []interface{}{[]interface{}{1}, []interface{}{2, 3}, []interface{}{4, 5}}

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

func TestLevelOrderTraverse(t *testing.T) {
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
		expected []interface{}
	}{
		{t1, []interface{}{[]interface{}{0}}},
		{t2, []interface{}{[]interface{}{1}}},
		{t3, []interface{}{[]interface{}{1}, []interface{}{2}, []interface{}{3}}},
		{t4, []interface{}{[]interface{}{1}, []interface{}{2, 3}}},
		{t5, []interface{}{[]interface{}{1}, []interface{}{2, 3}, []interface{}{4, 5}}},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			levelOrderTraverse(tt.in),
		)
	}
}

func levelOrderTraverse(root *common.TreeNode) []interface{} {
	out := common.NewList()

	if root == nil {
		return out.Slice()
	}

	// initialize a queue with the root.
	queue := common.NewQueue()
	queue.Push(root)

	for queue.Size() > 0 {
		levelSize := queue.Size()
		currentLevel := common.NewList()

		for i := 0; i < levelSize; i++ {
			// pop the queue and cache that value to its current level.
			current := queue.Pop().(*common.TreeNode)
			currentLevel.PushBack(current.Value)

			// push its left child.
			if current.Left != nil {
				queue.Push(current.Left)
			}

			// push its right child.
			if current.Right != nil {
				queue.Push(current.Right)
			}
		}

		out.PushBack(currentLevel.Slice())
	}

	return out.Slice()
}
