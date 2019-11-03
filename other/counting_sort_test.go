/*
Problem:
- Implement counting sort.

Approach:
- Iterate through the input, count the number of times each item occurs, use
  these counts to compute each item's index in the final sorted array.

Cost:
- O(n) time and O(n) space.
*/

package other

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestCountingSort(t *testing.T) {
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
		result := countingSort(tt.in, 6)
		common.Equal(t, tt.expected, result)
	}
}

func countingSort(in []int, max int) []int {
	// utilize max value to create a fix-sized list of item counts.
	counts := make([]int, max+1)
	out := make([]int, 0)

	// populate the array where its indices represent items themselves and
	// its values represent how many time the item appears.
	for _, item := range in {
		counts[item]++
	}

	// iterate through the counts and add the item to the output list.
	for i := 0; i < len(counts); i++ {
		count := counts[i]

		for j := 0; j < count; j++ {
			out = append(out, i)
		}
	}

	return out
}
