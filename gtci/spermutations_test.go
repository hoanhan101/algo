/*
Problem:
- Given a set of distinct numbers, find all of its permutations.

Example:
- Input: [1 2]
  Output: [[2 1] [1 2]]
- Input: [1 2 3]
  Output: [[3 2 1] [2 3 1] [2 1 3] [3 1 2] [1 3 2] [1 2 3]]

Approach:
- For every existing permutations, create new permutations by adding the number to
  all of its positions.

Cost:
- O(n*n!) time, O(n*n!) space since n! is the total of permutations of a n-number set.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFindSPermutations(t *testing.T) {
	tests := []struct {
		in       []int
		expected [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{1}, [][]int{{1}}},
		{[]int{1, 2}, [][]int{{2, 1}, {1, 2}}},
		{[]int{1, 2, 3}, [][]int{{3, 2, 1}, {2, 3, 1}, {2, 1, 3}, {3, 1, 2}, {1, 3, 2}, {1, 2, 3}}},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			findSPermutations(tt.in),
		)
	}
}

func findSPermutations(nums []int) [][]int {
	permutations := [][]int{}
	subsets := [][]int{}
	subsets = append(subsets, []int{})

	for _, num := range nums {
		n := len(subsets)
		for i := 0; i < n; i++ {
			current := subsets[i]
			// add the current number at every position of a new subset.
			for j := 0; j < len(current)+1; j++ {
				news := current

				news = append(news, 0)
				copy(news[j+1:], news[j:])
				news[j] = num

				// if the length of the new set is equal to the length of the
				// given slice, we found a permutation.
				if len(news) == len(nums) {
					permutations = append(permutations, news)
				}

				subsets = append(subsets, news)
			}
		}
	}

	return permutations
}
