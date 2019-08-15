// Problem:
// Implement DFS and BFS to traverse a tree.

package main

import (
	"reflect"
	"testing"
)

func TestTreeTraverse(t *testing.T) {
	tree := &Tree{nil, 1, nil}
	tree.left = &Tree{nil, 2, nil}
	tree.right = &Tree{nil, 3, nil}
	tree.left.left = &Tree{nil, 4, nil}
	tree.left.right = &Tree{nil, 5, nil}
	tree.right.right = &Tree{nil, 6, nil}

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

	r1 := []int{4, 2, 5, 1, 3, 6}
	e1 := chanToSlice(c1)
	if !reflect.DeepEqual(r1, e1) {
		t.Errorf("should be %v instead %v", e1, r1)
	}

	r2 := []int{1, 2, 4, 5, 3, 6}
	e2 := chanToSlice(c2)
	if !reflect.DeepEqual(r2, e2) {
		t.Errorf("should be %v instead %v", e2, r2)
	}

	r3 := []int{4, 5, 2, 6, 3, 1}
	e3 := chanToSlice(c3)
	if !reflect.DeepEqual(r3, e3) {
		t.Errorf("should be %v instead %v", e3, r3)
	}

	r4 := []int{1, 2, 3, 4, 5, 6}
	e4 := chanToSlice(c4)
	if !reflect.DeepEqual(r4, e4) {
		t.Errorf("should be %v instead %v", e4, r4)
	}
}

type Tree struct {
	left  *Tree
	value int
	right *Tree
}

// inorder DFS traverse.
func inorderTraverse(t *Tree, ch chan int) {
	if t == nil {
		return
	}

	inorderTraverse(t.left, ch)
	ch <- t.value
	inorderTraverse(t.right, ch)
}

// preorder DFS traverse.
func preorderTraverse(t *Tree, ch chan int) {
	if t == nil {
		return
	}

	ch <- t.value
	preorderTraverse(t.left, ch)
	preorderTraverse(t.right, ch)
}

// postorder DFS traverse.
func postorderTraverse(t *Tree, ch chan int) {
	if t == nil {
		return
	}

	postorderTraverse(t.left, ch)
	postorderTraverse(t.right, ch)
	ch <- t.value
}

// BFS traverse.
func levelorderTraverse(t *Tree, ch chan int) {
	if t == nil {
		return
	}

	// initialize a queue by enqueuing the root.
	queue := []*Tree{t}

	for len(queue) > 0 {
		// print front of queue and remove it.
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

func chanToSlice(ch chan int) []int {
	out := []int{}

	for v := range ch {
		out = append(out, v)
	}

	return out
}
