/*
Problem:
- Given a list of non-overlapping intervals sorted by their start time, insert
  a given interval at the correct position and merge all necessary intervals to
  produce a list that has only mutually exclusive intervals.

Example:
- Input: []interval{{1, 3}, {5, 7}, {8, 12}}, interval{4, 6}
  Output: []interval{{1, 3}, {4, 7}, {8, 12}}

Approach:
- Since the list is sorted, can skip all the intervals where their end time is
  less than the new interval's start time.
- The merging process is similar to the one in merge interval problem.

Cost:
- O(n) time, O(n) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestInsertInterval(t *testing.T) {
	tests := []struct {
		in1      []interval
		in2      interval
		expected []interval
	}{
		{[]interval{}, interval{}, []interval{}},
		{[]interval{{5, 7}, {8, 12}}, interval{1, 3}, []interval{{1, 3}, {5, 7}, {8, 12}}},
		{[]interval{{1, 3}, {5, 7}}, interval{8, 12}, []interval{{1, 3}, {5, 7}, {8, 12}}},
		{[]interval{{1, 3}, {5, 7}, {8, 12}}, interval{4, 6}, []interval{{1, 3}, {4, 7}, {8, 12}}},
		{[]interval{{1, 3}, {5, 7}, {8, 12}}, interval{3, 8}, []interval{{1, 12}}},
	}

	for _, tt := range tests {
		result := insertInterval(tt.in1, tt.in2)
		common.Equal(t, tt.expected, result)
	}
}

func insertInterval(intervals []interval, newInterval interval) []interval {
	if len(intervals) == 0 {
		return intervals
	}

	merged := []interval{}
	i := 0

	// skip all the intervals that comes before the new interval.
	for i < len(intervals) && intervals[i].end < newInterval.start {
		merged = append(merged, intervals[i])
		i++
	}

	// merge all the intervals that overlap with the new interval.
	for i < len(intervals) && intervals[i].start <= newInterval.end {
		newInterval.start = common.Min(intervals[i].start, newInterval.start)
		newInterval.end = common.Max(intervals[i].end, newInterval.end)
		i++
	}

	merged = append(merged, newInterval)

	// add all the remaining intervals to the output.
	for i < len(intervals) {
		merged = append(merged, intervals[i])
		i++
	}

	return merged
}
