/*
Problem:
- Given a sorted array, create a new array containing squares of all the
  number of the input array in the sorted order.

Assumption:
- The input can have negative numbers.

Example:
- Input: []int{-2, -1, 0, 1, 2}
  Output: []int{0, 1, 1, 4, 4}

Approach:
- The difficult part is to generate the output array with squares in sorted order because we have negative numbers and their squares are positive.
- Have one pointer start at the beginning and one at the end and let them
  move toward each other.
- At each step, whichever bigger will be added to the output array, from
  right to left.

Cost:
- O(n) time, O(n) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestSortedSquare(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{-3, -2, -1}, []int{1, 4, 9}},
		{[]int{1, 2, 3}, []int{1, 4, 9}},
		{[]int{-2, -1, 0, 1, 2}, []int{0, 1, 1, 4, 4}},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			sortedSquare(tt.in),
		)
	}
}

func sortedSquare(a []int) []int {
	out := make([]int, len(a))

	start := 0
	end := len(a) - 1
	outIndex := end

	for start <= end {
		startSquare := a[start] * a[start]
		endSquare := a[end] * a[end]

		if startSquare < endSquare {
			out[outIndex] = endSquare
			end--
		} else if startSquare > endSquare {
			out[outIndex] = startSquare
			start++
		} else {
			// only need to move one side if their squares are equal.
			out[outIndex] = startSquare
			start++
		}

		outIndex--
	}

	return out
}
