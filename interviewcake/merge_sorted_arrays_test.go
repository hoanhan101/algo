/*
Problem:
- Merge two sorted arrays.

Example:
- Input: []int{1, 3, 5}, []int{2, 4, 6}
  Output: []int{1, 2, 3, 4, 5, 6}
- Input: []int{1, 3, 5}, []int{2, 4, 6, 7}
  Output: []int{1, 2, 3, 4, 5, 6, 7},

Approach:
- Since these arrays are sorted, can use two pointers approach to iterate
  through both of them and append the smaller value to a new merged list at
  each step.

Solution:
- Have two pointers start at the beginning of these two arrays.
- While both of them does not reach the end, compare two current values
  at each step and append the smaller one two a new merged list.
- Move the two pointers up accordingly as values get merged in.
- In the case where one of these pointers reach the end first and the
  other one is still in the middle of the array, simply add the rest of
  its values to the merged list since they are all sorted and guaranteed
  to be in ascending order.

Cost:
- O(n) time, O(n) space.
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      []int
		expected []int
	}{
		{
			[]int{},
			[]int{},
			[]int{},
		},
		{
			[]int{},
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 3, 5},
			[]int{2, 4, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1, 3, 5},
			[]int{2, 4, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{6},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		result := mergeSortedArray(tt.in1, tt.in2)
		common.Equal(t, tt.expected, result)
	}
}

func mergeSortedArray(a1, a2 []int) []int {
	out := []int{}

	// keep two "pointers" at index 0 and move up accordingly as one get
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
