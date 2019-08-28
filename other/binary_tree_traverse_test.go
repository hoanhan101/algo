// Problem:
// Implement binary tree's depth first search (inorder, preorder, postorder)
// and breath-first search (levelorder).
//
// The solution uses a channel to send value over as we traverse the tree.

package main

import (
	"reflect"
	"testing"
)

func TestBinaryTreeTraverse(t *testing.T) {
	// initialize a test tree.
	tree := &BinaryTree{nil, 1, nil}
	tree.left = &BinaryTree{nil, 2, nil}
	tree.right = &BinaryTree{nil, 3, nil}
	tree.left.left = &BinaryTree{nil, 4, nil}
	tree.left.right = &BinaryTree{nil, 5, nil}
	tree.right.right = &BinaryTree{nil, 6, nil}

	// use 4 different channels for 4 different methods.
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	c4 := make(chan int)

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

	// define test cases accordingly.
	tests := []struct {
		c        chan int
		expected []int
	}{
		{c1, []int{4, 2, 5, 1, 3, 6}},
		{c2, []int{1, 2, 4, 5, 3, 6}},
		{c3, []int{4, 5, 2, 6, 3, 1}},
		{c4, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		result := chanToSlice(tt.c)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
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
	queue := []*BinaryTree{t}

	for len(queue) > 0 {
		// send the front of the queue and dequeue it.
		current := queue[0]
		ch <- current.value
		queue = queue[1:]

		// enqueue its left child.
		if current.left != nil {
			queue = append(queue, current.left)
		}

		// enqueue its right child.
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
}

// chanToSlice pushes values from a channel to a slice.
func chanToSlice(ch chan int) []int {
	out := []int{}

	for v := range ch {
		out = append(out, v)
	}

	return out
}
