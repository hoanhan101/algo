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

// MaxHeap represents a max heap data structure.
type MaxHeap struct {
	heap *maxHeap
}

// NewMaxHeap returns a new max heap.
func NewMaxHeap() *MaxHeap {
	h := &maxHeap{}
	heap.Init(h)
	return &MaxHeap{heap: h}
}

// Push adds a new node to the heap.
func (h *MaxHeap) Push(x int) {
	heap.Push(h.heap, x)
}

// Pop returns the root node after removing it from the heap.
func (h *MaxHeap) Pop() int {
	return heap.Pop(h.heap).(int)
}

// Len returns the number of items in the heap.
func (h *MaxHeap) Len() int {
	return h.heap.Len()
}

// Peek returns the maximum item in the heap.
func (h *MaxHeap) Peek() int {
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

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
