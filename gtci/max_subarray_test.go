/*
Problem:
- Given an array of positive numbers and a positive number k, find the maximum
  sum of any contiguous subarray of size k.

Example:
- Input: int{2, 1, 5, 1, 3, 2}, k=3
  Output: 9
  Explanation: Subarray with maximum sum is [5, 1, 3].

Approach:
- View each contiguous subarray as a sliding window of k elements.
- As we move to the next subarray, slide the window by one element to
  reuse the sum from previous array and update the maximum sum.

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestMaxSubarray(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      int
		expected int
	}{
		{[]int{}, 1, 0},
		{[]int{1}, 1, 1},
		{[]int{1}, 2, 0},
		{[]int{1, 2}, 1, 2},
		{[]int{1, 2}, 2, 3},
		{[]int{1, 2}, 3, 0},
		{[]int{2, 1, 5, 1, 3, 2}, 3, 9},
		{[]int{2, 3, 4, 1, 5}, 2, 7},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			maxSubarray(tt.in1, tt.in2),
		)
	}
}

func maxSubarray(a []int, k int) int {
	max := 0

	// sum keeps track of the sum of a window while start keeps track of
	// its start index.
	sum, start := 0, 0

	for end := range a {
		sum += a[end]

		// slide the window once we hit its limit.
		if end >= k-1 {
			max = common.Max(max, sum)

			// subtract the start element and increase the start index to move
			// the window ahead by one element.
			sum -= a[start]
			start++
		}
	}

	return max
}
