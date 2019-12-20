/*
Problem:
- Given two sorted lists of intervals, find the intersection between them.

Example:
- Input: []interval{{1, 3}, {5, 6}, {7, 9}}, []interval{{2, 3}, {5, 7}}
  Output: []interval{{2, 3}, {5, 6}, {7, 7}}

Approach:
- Iterate through both list at the same time and check if two intervals
  are overlapping at each step.
- Opposite to the merging process, an overlapped interval has:
  - a bigger start between the two intervals
  - a smaller end between the two intervals

Cost:
- O(n) time, O(n) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestIntersectIntervals(t *testing.T) {
	tests := []struct {
		in1      []interval
		in2      []interval
		expected []interval
	}{
		{[]interval{}, []interval{}, []interval{}},
		{[]interval{{1, 3}, {5, 6}, {7, 9}}, []interval{{2, 3}, {5, 7}}, []interval{{2, 3}, {5, 6}, {7, 7}}},
		{[]interval{{1, 3}, {5, 7}, {9, 12}}, []interval{{5, 10}}, []interval{{5, 7}, {9, 10}}},
	}

	for _, tt := range tests {
		result := intersectIntervals(tt.in1, tt.in2)
		common.Equal(t, tt.expected, result)
	}
}

func intersectIntervals(intervals1, intervals2 []interval) []interval {
	out := []interval{}
	i, j := 0, 0

	for i < len(intervals1) && j < len(intervals2) {
		// if two intervals are overlapping, insert an interval that has a
		// bigger start between the two intervals and a smaller end between them.
		if overlap(intervals1[i], intervals2[j]) || overlap(intervals2[j], intervals1[i]) {
			out = append(
				out,
				interval{
					common.Max(intervals1[i].start, intervals2[j].start),
					common.Min(intervals1[i].end, intervals2[j].end),
				},
			)
		}

		// move the pointer up accordingly.
		if intervals1[i].end < intervals2[j].end {
			i++
		} else {
			j++
		}
	}

	return out
}

func overlap(in1, in2 interval) bool {
	return in1.start >= in2.start && in1.start <= in2.end
}
