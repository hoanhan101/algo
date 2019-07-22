package main

import (
	"reflect"
	"sort"
	"testing"
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
	}

	for _, tt := range tests {
		result := mergeMeetings(tt.in)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

type meeting struct {
	start int
	end   int
}

func mergeMeetings(meetings []meeting) []meeting {
	// sort the meetings in ascending order
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].start < meetings[j].start
	})

	out := []meeting{}
	for i := range meetings {
		// push the first one to the list so we can have a start
		if i == 0 {
			out = append(out, meetings[i])
			continue
		}

		// check against the last merged one.
		if out[len(out)-1].end >= meetings[i].start {
			out[len(out)-1].end = larger(meetings[i].end, out[len(out)-1].end)
		} else {
			out = append(out, meetings[i])
		}
	}

	return out
}

func larger(a, b int) int {
	if a > b {
		return a
	}

	return b
}
