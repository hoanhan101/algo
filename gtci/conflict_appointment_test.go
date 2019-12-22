/*
Problem:
- Given a list of intervals, check if any of them is conflicting.

Example:
- Input: []interval{{1, 2}, {2, 3}, {4, 5}}
  Output: false
- Input: []interval{{1, 5}, {2, 3}}
  Output: true

Approach:
- Similar to merge intervals problem, need to return the true
  immediately if any of them is conflicting.

Cost:
- O(nlogn) time, O(n) space.
*/

package gtci

import (
	"sort"
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestIsConflicting(t *testing.T) {
	tests := []struct {
		in       []interval
		expected bool
	}{
		{[]interval{}, false},
		{[]interval{{1, 2}}, false},
		{[]interval{{1, 2}, {2, 3}}, false},
		{[]interval{{2, 3}, {1, 2}}, false},
		{[]interval{{1, 2}, {2, 3}, {5, 6}}, false},
		{[]interval{{1, 2}, {5, 6}, {2, 3}}, false},
		{[]interval{{1, 2}, {1, 3}}, true},
		{[]interval{{1, 3}, {1, 3}}, true},
		{[]interval{{1, 2}, {2, 3}, {2, 6}}, true},
		{[]interval{{1, 2}, {2, 6}, {2, 3}}, true},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			isConflicting(tt.in),
		)
	}
}

func isConflicting(intervals []interval) bool {
	// sort the intervals in ascending order.
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	// return false immediately if any interval is conflicting.
	for i := 1; i < len(intervals); i++ {
		if intervals[i].start < intervals[i-1].end {
			return true
		}
	}

	return false
}
