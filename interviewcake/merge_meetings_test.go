// Problem:
// Given a list of unsorted, independent meetings, returns a list of a merged
// one.
//
// Example:
// Given:  []meeting{{1, 2}, {2, 3}, {4, 5}}
// Return: []meeting{{1, 3}, {4, 5}}
//
// Solution:
// Sort the list in ascending order so that meetings that might need to be
// merged are next to each other.
// Iterate through the list and check if the last merged meeting end time
// is greater or equal then the current one start time, merge them using the
// later end time.
//
// Cost:
// O(nlogn) time, O(n) space.
// Because we sort all meeting first, the runtime is O(nlogn). We create a new
// list of merged meeting times, so the space cost is O(n).

package interviewcake

import (
	"sort"
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestMergeMeetings(t *testing.T) {
	tests := []struct {
		in       []meeting
		expected []meeting
	}{
		{[]meeting{}, []meeting{}},
		{[]meeting{{1, 2}}, []meeting{{1, 2}}},
		{[]meeting{{1, 2}, {2, 3}}, []meeting{{1, 3}}},
		{[]meeting{{1, 2}, {4, 5}}, []meeting{{1, 2}, {4, 5}}},
		{[]meeting{{1, 2}, {2, 3}, {4, 5}}, []meeting{{1, 3}, {4, 5}}},
		{[]meeting{{1, 6}, {2, 3}, {4, 5}}, []meeting{{1, 6}}},
		{[]meeting{{4, 5}, {2, 3}, {1, 6}}, []meeting{{1, 6}}},
	}

	for _, tt := range tests {
		result := mergeMeetings(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

// meeting has a start and end time.
type meeting struct {
	start int
	end   int
}

func mergeMeetings(meetings []meeting) []meeting {
	// sort the meetings in ascending order.
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].start < meetings[j].start
	})

	out := []meeting{}
	for i := range meetings {
		// push the first meeting to the list so we can have a start.
		if i == 0 {
			out = append(out, meetings[i])
			continue
		}

		// if the last merged meeting's end time is greater or equal than the current
		// one's start time, merge them by using the later ending time. else,
		// leave them separate and push it to the output list.
		if out[len(out)-1].end >= meetings[i].start {
			_, out[len(out)-1].end = common.Mimax(meetings[i].end, out[len(out)-1].end)
		} else {
			out = append(out, meetings[i])
		}
	}

	return out
}
