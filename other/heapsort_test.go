/*
Problem:
- Implement heapsort.

Approach:
- Similar to selection sort, repeatedly choose the largest item and move it to
  the end of the array using a max heap.

Cost:
- O(nlogn) time and O(1) space.
*/

package other

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestHeapsort(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2, 1, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
		{[]int{2, 1, 2}, []int{1, 2, 2}},
		{[]int{1, 2, 4, 3, 6, 5}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{6, 2, 4, 3, 1, 5}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		heapsort(tt.in)
		common.Equal(t, tt.expected, tt.in)
	}
}

func heapsort(in []int) {
	heapify(in)

	size := len(in)
	for size > 0 {
		// repeatedly remove the largest item.
		largest := removeLargest(in, size)

		// update the heap size.
		size--

		// store the removed value at the end of the list.
		in[size] = largest
	}
}

// heapify transform the input into a max heap.
func heapify(in []int) {
	for i := len(in) - 1; i > -1; i-- {
		bubbleDown(in, len(in), i)
	}
}

// bubbleDown allow larger values to reach the top.
func bubbleDown(heap []int, heapSize int, index int) {
	for index < heapSize {
		// fast-calculate the children left and right index.
		left := index*2 + 1
		right := index*2 + 2

		// stop if there is no child node.
		if left >= heapSize {
			break
		}

		// find the larger index
		larger := left
		if right < heapSize && heap[left] < heap[right] {
			larger = right
		}

		// if the current item is larger than both children, we're done.
		// if not, swap with the larger child.
		if heap[index] < heap[larger] {
			common.Swap(heap, index, larger)
		} else {
			break
		}
	}
}

// removeLargest remove and return the largest item from the heap.
func removeLargest(heap []int, heapSize int) int {
	// largest item is at the top of our max heap.
	largest := heap[0]

	// move the last item into the root position.
	heap[0] = heap[heapSize-1]

	// bubble down from the root to restore the heap.
	bubbleDown(heap, heapSize-1, 0)

	return largest
}
