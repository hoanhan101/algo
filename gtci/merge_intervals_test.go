/*
Problem:
- Given a list of intervals, merge all the overlapping intervals to produce
  a list that has only mutually exclusive intervals.

Example:
- Input: []interval{{1, 2}, {2, 3}, {4, 5}}
  Output: []interval{{1, 3}, {4, 5}}
- Input: []interval{{1, 5}, {2, 3}}
  Output: []interval{{1, 5}}

Approach:
- Sort the list in ascending order so that intervals that might need to be
  merged are next to each other.
- Can merge two intervals together if the first one's end time is greater or
  or equal than the second one's start time.

Cost:
- O(nlogn) time, O(n) space.
- Because we sort all intervals first, the runtime is O(nlogn). We create a new
  list of merged interval times, so the space cost is O(n).
*/

package gtci

import (
	"sort"
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		in       []interval
		expected []interval
	}{
		{[]interval{}, []interval{}},
		{[]interval{{1, 2}}, []interval{{1, 2}}},
		{[]interval{{1, 2}, {2, 3}}, []interval{{1, 3}}},
		{[]interval{{1, 5}, {2, 3}}, []interval{{1, 5}}},
		{[]interval{{1, 2}, {4, 5}}, []interval{{1, 2}, {4, 5}}},
		{[]interval{{1, 5}, {2, 3}, {4, 5}}, []interval{{1, 5}}},
		{[]interval{{1, 2}, {2, 3}, {4, 5}}, []interval{{1, 3}, {4, 5}}},
		{[]interval{{1, 6}, {2, 3}, {4, 5}}, []interval{{1, 6}}},
		{[]interval{{4, 5}, {2, 3}, {1, 6}}, []interval{{1, 6}}},
	}

	for _, tt := range tests {
		result := mergeIntervals(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

// interval has a start and end time.
type interval struct {
	start int
	end   int
}

func mergeIntervals(intervals []interval) []interval {
	// sort the intervals in ascending order.
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	merged := []interval{}
	for i := range intervals {
		// push the first interval to the list so we can have a start.
		if i == 0 {
			merged = append(merged, intervals[i])
			continue
		}

		// if the last merged interval's end time is greater or equal than the current
		// one's start time, merge them by using the larger ending time. else,
		// leave them separate and push it to the merged list.
		if merged[len(merged)-1].end >= intervals[i].start {
			merged[len(merged)-1].end = common.Max(intervals[i].end, merged[len(merged)-1].end)
		} else {
			merged = append(merged, intervals[i])
		}
	}

	return merged
}
