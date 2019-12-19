/*
Problem:
- Given an array containing 0s and 1s, if you are allowed to replace no more
  than k 0s with 1s, find the length of the longest contiguous subarray having all 1s.

Example:
- Input: []int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}, k=2
  Output: 6
  Explanation: Have the longest subarray of 1s after replacing 0 at index 5 and 8

Approach:
- Similar to longest substring after k replacements problem, except we only have
  1 and 0 in the array.

Cost:
- O(n) time, O(1) space since there are only 26 characters in the alphabet.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestLongestSubstringOnesReplacement(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      int
		expected int
	}{
		{[]int{}, 1, 0},
		{[]int{0}, 1, 1},
		{[]int{0, 0}, 1, 1},
		{[]int{0, 0, 1}, 1, 2},
		{[]int{1, 0, 1}, 1, 3},
		{[]int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}, 2, 6},
		{[]int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1}, 3, 9},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			longestSubstringOnesReplacement(tt.in1, tt.in2),
		)
	}
}

func longestSubstringOnesReplacement(arr []int, k int) int {
	maxLength, start, maxOnesCount := 0, 0, 0

	for end := range arr {
		if arr[end] == 1 {
			maxOnesCount++
		}

		// shrink ther window as we are not allowed to replace more than k 0s.
		if end-start+1-maxOnesCount > k {
			if arr[start] == 1 {
				maxOnesCount--
			}

			start++
		}

		// update the maximum length at each step.
		maxLength = common.Max(maxLength, end-start+1)
	}

	return maxLength
}
