/*
Problem:
- Implement binary tree's depth first search (inorder, preorder, postorder)
  and breath-first search (levelorder).

Approach:
- The solution uses a channel to send value over as we traverse the tree.
*/

package other

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestBinaryTreeTraverse(t *testing.T) {
	// predefine a test tree that has the following structure.
	//         1
	//      2     3
	//    4   5     6
	tree := &BinaryTree{nil, 1, nil}
	tree.left = &BinaryTree{nil, 2, nil}
	tree.right = &BinaryTree{nil, 3, nil}
	tree.left.left = &BinaryTree{nil, 4, nil}
	tree.left.right = &BinaryTree{nil, 5, nil}
	tree.right.right = &BinaryTree{nil, 6, nil}

	// use 4 different channels for 4 different methods to send and receive
	// values as we traverse the tree.
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	c4 := make(chan int)

	// close the channels after sending values.
	go func() {
		inorderTraverse(tree, c1)
		close(c1)
	}()

	go func() {
		preorderTraverse(tree, c2)
		close(c2)
	}()

	go func() {
		postorderTraverse(tree, c3)
		close(c3)
	}()

	go func() {
		levelorderTraverse(tree, c4)
		close(c4)
	}()

	// push the values from channels to slices and define test cases against
	// them.
	tests := []struct {
		c        chan int
		expected []int
	}{
		{c1, []int{4, 2, 5, 1, 3, 6}}, // inorder
		{c2, []int{1, 2, 4, 5, 3, 6}}, // preorder
		{c3, []int{4, 5, 2, 6, 3, 1}}, // postorder
		{c4, []int{1, 2, 3, 4, 5, 6}}, // breath-search, aka levelorder
	}

	for _, tt := range tests {
		result := common.ChanToSlice(tt.c)
		common.Equal(t, tt.expected, result)
	}
}

type BinaryTree struct {
	left  *BinaryTree
	value int
	right *BinaryTree
}

// inorder DFS traverse.
func inorderTraverse(t *BinaryTree, ch chan int) {
	if t == nil {
		return
	}

	inorderTraverse(t.left, ch)
	ch <- t.value
	inorderTraverse(t.right, ch)
}

// preorder DFS traverse.
func preorderTraverse(t *BinaryTree, ch chan int) {
	if t == nil {
		return
	}

	ch <- t.value
	preorderTraverse(t.left, ch)
	preorderTraverse(t.right, ch)
}

// postorder DFS traverse.
func postorderTraverse(t *BinaryTree, ch chan int) {
	if t == nil {
		return
	}

	postorderTraverse(t.left, ch)
	postorderTraverse(t.right, ch)
	ch <- t.value
}

// levelorder BFS traverse.
func levelorderTraverse(t *BinaryTree, ch chan int) {
	if t == nil {
		return
	}

	// initialize a queue by enqueuing the root.
	queue := common.NewQueue()
	queue.Push(t)

	for queue.Size() > 0 {
		// pop the queue and send that value to the channel.
		current := queue.Pop().(*BinaryTree)
		ch <- current.value

		// push its left child.
		if current.left != nil {
			queue.Push(current.left)
		}

		// push its right child.
		if current.right != nil {
			queue.Push(current.right)
		}
	}
}
