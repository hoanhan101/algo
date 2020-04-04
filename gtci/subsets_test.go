/*
Problem:
- Given a set with distinct elements, find all the distinct subsets.

Example:
- Input: [1 2]
  Output: [[] [1] [2] [1 2]]
- Input: [1 2 5]
  Output: [[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]

Approach:
- Start with an empty set.
- Iterate through the set one by one and add them to existing sets to create new ones.

Cost:
- O(2^n) time, O(2^n) space since we would have a total of 2^n subsets.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFindSubsets(t *testing.T) {
	tests := []struct {
		in       []int
		expected [][]int
	}{
		{[]int{}, [][]int{{}}},
		{[]int{1}, [][]int{{}, {1}}},
		{[]int{1, 2}, [][]int{{}, {1}, {2}, {1, 2}}},
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			findSubsets(tt.in),
		)
	}
}

func findSubsets(nums []int) [][]int {
	subsets := [][]int{}

	// start with an empty set.
	subsets = append(subsets, []int{})

	// for each number in the set, add it to all existing sets.
	for _, num := range nums {
		n := len(subsets)
		for i := 0; i < n; i++ {
			set := subsets[i]
			set = append(set, num)
			subsets = append(subsets, set)
		}
	}

	return subsets
}
