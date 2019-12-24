/*
Problem:
- Implement bubble sort.

Approach:
- Repeatedly swap the adjacent elements if they are in the wrong order in the
  array, one item at a time.

Cost:
- O(n^2) time and O(1) space.
*/

package other

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestBubbleSort(t *testing.T) {
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
		bubbleSort(tt.in)
		common.Equal(t, tt.expected, tt.in)
	}
}

func bubbleSort(in []int) {
	length := len(in)

	// for each element in the list, check it with almost every other element.
	for i := 0; i < length; i++ {
		// since the last i element is already in place, only iterate through
		// the item before the last one.
		for j := 0; j < length-i-1; j++ {
			// swap the adjacent elements if they are not in ascending order.
			if in[j] > in[j+1] {
				common.Swap(in, j, j+1)
			}
		}
	}
}
