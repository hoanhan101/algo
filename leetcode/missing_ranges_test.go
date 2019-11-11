/*
Problem:
- Given a sorted integer array where the range of elements are [0, 99] inclusive,
  return its missing ranges.

Example:
- Input: []int{0, 1, 6, 16, 66, 99}
  Output: []string{"2-5", "7-15", "17-65", "67-98"}
- Input: []int{6, 16, 66}
  Output: []string{"0-5", "7-15", "17-65", "67-99"}

Approach:
- Keep two pointers where one is ahead of the other by 1 index.
- Iterate through the list, calculate the difference of two consecutive numbers
  in the list at each step and append it to a new list.

Solution:
- Initialize a list of string as an output, two pointers where one is ahead of
  the other by 1.
- While p2 is less than the length of the input list, check if the difference
  between two current numbers are greater or equal to 2.
- If so, add it.
- Otherwise, no need to add it because they are consecutive numbers already so
  there is no range to get.
- Make sure to check the edge case where the first value in the input list is
  either 0 or 99.

Cost:
- O(n) time, O(m) space, where m < n and n is the size of the input.
*/

package leetcode

import (
	"fmt"
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFindMissingRanges(t *testing.T) {
	tests := []struct {
		in       []int
		expected []string
	}{
		{[]int{}, []string{}},
		{[]int{6}, []string{"0-5", "7-99"}},
		{[]int{0, 6}, []string{"1-5", "7-99"}},
		{[]int{6, 99}, []string{"0-5", "7-98"}},
		{[]int{0, 6, 99}, []string{"1-5", "7-98"}},
		{[]int{6, 16, 66}, []string{"0-5", "7-15", "17-65", "67-99"}},
		{[]int{0, 1, 6, 16, 66, 99}, []string{"2-5", "7-15", "17-65", "67-98"}},
	}

	for _, tt := range tests {
		result := findMissingRanges(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func TestGetRange(t *testing.T) {
	tests := []struct {
		in1      int
		in2      int
		expected string
	}{
		{-1, 100, "0-99"},
		{0, 99, "1-98"},
		{1, 98, "2-97"},
		{6, 66, "7-65"},
	}

	for _, tt := range tests {
		result := getRange(tt.in1, tt.in2)
		common.Equal(t, tt.expected, result)
	}
}

func findMissingRanges(list []int) []string {
	// initialize a list of string as an output.
	out := []string{}

	p1 := 0
	p2 := p1 + 1

	// check edge case where the list is empty.
	if len(list) == 0 {
		return out
	}

	// check edge case where the first value is not 0.
	if list[0] != 0 {
		out = append(out, getRange(-1, list[0]))
	}

	for p2 < len(list) {
		// if the difference between two numbers are greater or equal to 2 then
		// add it. otherwise, no need to add because they are consecutive
		// numbers already so there is no range to get.
		if list[p2]-list[p1] >= 2 {
			out = append(out, getRange(list[p1], list[p2]))
		}

		p1++
		p2++
	}

	// check edge case where the last value is not 99.
	if list[len(list)-1] != 99 {
		out = append(out, getRange(list[len(list)-1], 100))
	}

	return out
}

// getRange returns the string representation of the range between the lower
// bound and upper bound.
func getRange(start, end int) string {
	return fmt.Sprintf("%v-%v", start+1, end-1)
}
