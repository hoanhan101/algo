/*
Problem:
- Given an array of positive numbers and a positive number s, find the length
  of the smallest contiguous subarray whose sum is greater than or equal to s.

Example:
- Input: array=int{2, 1, 5, 2, 3, 2}, s=7
  Output: 2
  Explanation: Smallest subarray with a sum great than or equal to 7 is [5, 2]
  with length=2.

Approach:
- The difference between the previous problem and this one is that the size of
  the sliding window is not fixed.
- Can still use the similar strategy to add up elements until their sum is greater
  than equal to s and view them as our sliding window.
- Shrink the window until the window's sum is smaller than s again while keep
  updating the minimum length.

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"math"
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestMinSubarray(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      int
		expected int
	}{
		{[]int{}, 1, 0},
		{[]int{2, 1, 5, 2, 3, 2}, 7, 2},
		{[]int{2, 1, 5, 2, 3, 2}, 8, 3},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			minSubarray(tt.in1, tt.in2),
		)
	}
}

func minSubarray(a []int, s int) int {
	minLength := math.MaxInt64

	// sum keeps track of the sum of a window while start keeps track of
	// its start index.
	sum, start := 0, 0

	for end := range a {
		sum += a[end]

		// shrink the window until the window's sum is smaller than the
		// target sum.
		for sum >= s {
			// update the minimum length at each step.
			minLength = common.Min(minLength, end-start+1)

			// subtract the start element and increase the start index to move
			// the window ahead by one element.
			sum -= a[start]
			start++
		}
	}

	// let min=0 if there is no such subarray exists.
	if minLength == math.MaxInt64 {
		minLength = 0
	}

	return minLength
}
