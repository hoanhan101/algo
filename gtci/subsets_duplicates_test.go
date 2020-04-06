/*
Problem:
- Given a set with duplicate elements, find all the distinct subsets.

Example:
- Input: [1 2]
  Output: [[] [1] [2] [1 2]]
- Input: [1 2 5]
  Output: [[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]

Approach:
- Similar to the previous problem, but we do two more steps.
- First, sort the set so that duplicates are next to each other.
- Second, when we encounter a duplicate while iterating the set,
  only create subsets from the subsets that added previously.
- Can use a two-pointer approach to update their start and end window
  accordingly.

Cost:
- O(2^n) time, O(2^n) space since we would have a total of 2^n subsets.
*/

package gtci

import (
	"sort"
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFindSubsetsWithDuplicates(t *testing.T) {
	tests := []struct {
		in       []int
		expected [][]int
	}{
		{[]int{}, [][]int{{}}},
		{[]int{1}, [][]int{{}, {1}}},
		{[]int{1, 2}, [][]int{{}, {1}, {2}, {1, 2}}},
		{[]int{1, 2, 2}, [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}}},
		{[]int{2, 1, 2}, [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}}},
		{[]int{2, 2, 1}, [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}}},
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{[]int{3, 2, 1}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			findSubsetsWithDuplicates(tt.in),
		)
	}
}

func findSubsetsWithDuplicates(nums []int) [][]int {
	// sort the set in ascending order.
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// start with an empty set.
	subsets := [][]int{}
	subsets = append(subsets, []int{})

	start, end := 0, 0

	// for each number in the set, add it to all existing sets.
	for i := 0; i < len(nums); i++ {
		// if the current item and the previous one are the same, update the
		// start index accordingly so that we only create subset from the one
		// that added in the previous step.
		start = 0
		if i > 0 && nums[i] == nums[i-1] {
			start = end + 1
		}

		end = len(subsets) - 1

		for j := start; j < end+1; j++ {
			set := subsets[j]
			set = append(set, nums[i])
			subsets = append(subsets, set)
		}
	}

	return subsets
}
