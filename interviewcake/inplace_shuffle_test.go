/*
Problem:
- Given a list of integers, shuffle its location in-place.

Example:
- Input: []int{1, 2, 3, 4, 5}
  Output: []int{2, 1, 4, 3, 5}

Approach:
- Iterate through the list, swap current value with a value in a randomized
  index that is between the current and last index.

Solution:
- Cache the last index value of the list.
- Iterate through the list, get a randomized index value between the
  current index and the last index, then use it to swap the corresponding
  values at these two indices.

Cost:
- O(n) time, O(1) space.
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestInplaceShuffle(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{2, 1, 4, 3, 5}},
	}

	for _, tt := range tests {
		result := inplaceShuffle(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func inplaceShuffle(list []int) []int {
	// check edge case.
	if len(list) <= 1 {
		return list
	}

	lastIndex := len(list) - 1
	for i := 0; i < len(list); i++ {
		// get a andomized index that is between the current and last index.
		randomIndex := common.Random(i, lastIndex)

		// swap current value.
		if i != randomIndex {
			common.Swap(list, i, randomIndex)
		}
	}

	return list
}
