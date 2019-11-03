/*
Problem:
- Implement merge sort.

Approach:
- Split the input in half, recursively sorts each half, then merge the
  sorted halves back together.

Cost:
- O(nlogn) time and O(n) space.
*/

package other

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestMergeSort(t *testing.T) {
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
		result := mergeSort(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func mergeSort(in []int) []int {
	// base case
	if len(in) <= 1 {
		return in
	}

	// split the input in half.
	middleIndex := len(in) / 2
	left := in[:middleIndex]
	right := in[middleIndex:]

	// sort each half.
	leftSorted := mergeSort(left)
	rightSorted := mergeSort(right)

	// merge the sorted halves.
	return mergeSortedArray(leftSorted, rightSorted)
}

func mergeSortedArray(a1, a2 []int) []int {
	out := []int{}

	// keep two "pointer" at index 0 and move up accordingly as one get
	// merged in.
	i, j := 0, 0
	for i < len(a1) && j < len(a2) {
		if a1[i] < a2[j] {
			out = append(out, a1[i])
			i++
		} else {
			out = append(out, a2[j])
			j++
		}
	}

	// if we get here, one array must have bigger size than the other. could
	// figure out which one is it then copy the rest of its to our final one.
	if i < len(a1) {
		out = append(out, a1[i:]...)
	}

	if j < len(a2) {
		out = append(out, a2[j:]...)
	}

	return out
}
