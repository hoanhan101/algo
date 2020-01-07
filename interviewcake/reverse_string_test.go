/*
Problem:
- Given a list of string, reverse its order in place.

Example:
- Input: []string{"a", "b", "c", "d"}
  Output: []string{"d", "c", "b", "a"}

Approach:
- Use two pointers approach to swap two values on both ends as we move toward
  the middle.

Solution:
- Initialize the two pointers, one starts at the beginning and one starts at
  the end of the list.
- While the start pointer does not meet the end pointer in the middle, keep
  swapping these two values.
- Move the start pointer up and move the end pointer down.

Cost:
- O(n) time, O(1) space.
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		in       []string
		expected []string
	}{
		{[]string{}, []string{}},
		{[]string{"a"}, []string{"a"}},
		{[]string{"a", "b"}, []string{"b", "a"}},
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}},
		{[]string{"a", "b", "c", "d"}, []string{"d", "c", "b", "a"}},
	}

	for _, tt := range tests {
		result := reverseString(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func reverseString(list []string) []string {
	// initialize start and end index pointer.
	start := 0
	end := len(list) - 1

	for start < end {
		// swap 2 character using a temp variable.
		common.Swap(list, start, end)

		// move the cursor toward the middle.
		start++
		end--
	}

	return list
}
