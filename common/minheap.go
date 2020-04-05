package common

import (
	"container/heap"
)

// MinHeap represents a min heap data structure.
type MinHeap struct {
	heap *minHeap
}

// NewMinHeap returns a new min heap.
func NewMinHeap() *MinHeap {
	h := &minHeap{}
	heap.Init(h)
	return &MinHeap{heap: h}
}

// Push adds a new node to the heap.
func (h *MinHeap) Push(x int) {
	heap.Push(h.heap, x)
}

// Pop returns the root node after removing it from the heap.
func (h *MinHeap) Pop() int {
	return heap.Pop(h.heap).(int)
}

// Len returns the number of items in the heap.
func (h *MinHeap) Len() int {
	return h.heap.Len()
}

// Peek returns the minimum item in the heap.
func (h *MinHeap) Peek() int {
	return (*h.heap)[0]
}

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
