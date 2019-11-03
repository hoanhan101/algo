/*
Problem:
- Implement insertion sort.

Approach:
- Insert elements from an unsorted array into a sorted subsection of the
  array, one item at a time.

Cost:
- O(n^2) time and O(1) space.
*/

package other

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestInsertionSort(t *testing.T) {
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
		insertionSort(tt.in)
		common.Equal(t, tt.expected, tt.in)
	}
}

func insertionSort(in []int) {
	// iterate through the list from position 1.
	for i := 1; i < len(in); i++ {
		// shift each one to the left by swapping it with the one before until
		// it's in the right spot.
		current := in[i]
		j := i - 1

		for j >= 0 && current < in[j] {
			in[j+1] = in[j]
			j--
		}

		in[j+1] = current
	}
}
