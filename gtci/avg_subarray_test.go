/*
Problem:
- Given an array, find the average of all contiguous subarrays of size k.

Example:
- Input: []int{1, 3, 2, 6, -1, 4, 1, 8, 2}, k=5
  Output: []int{2.2, 2.8, 2.4, 3.6, 2.8}

Approach:
- View each contiguous subarray as a sliding window of k elements.
- As we move to the next subarray, slide the window by one element to
  reuse the sum from previous array.

Cost:
- O(n) time, O(k) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestAvgSubarray(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      int
		expected []float64
	}{
		{[]int{}, 5, []float64{}},
		{[]int{1, 3, 2, 6, -1}, 5, []float64{2.2}},
		{[]int{1, 3, 2, 6, -1, 4, 1, 8, 2}, 5, []float64{2.2, 2.8, 2.4, 3.6, 2.8}},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			avgSubarray(tt.in1, tt.in2),
		)
	}
}

func avgSubarray(a []int, k int) []float64 {
	out := []float64{}

	// sum keeps track of the sum of a window while start keeps track of
	// its start index.
	sum, start := 0, 0

	for end := range a {
		sum += a[end]

		// slide the window once we hit its limit.
		if end >= k-1 {
			out = append(out, float64(sum)/float64(k))

			// subtract the start element and increase the start index to move
			// the window ahead by one element.
			sum -= a[start]
			start++
		}
	}

	return out
}
