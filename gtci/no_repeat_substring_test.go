/*
Problem:
- Given a string, find the length of the longest substring which has no repeating characters.

Example:
- Input: "aabccbb"
  Output: 3
  Explanation: Longest substring which has no repeating characters.is "abc"

Approach:
- Similar to longest substring with k distinct characters, we can use a hashmap
  to remember the last index of each character we have processed.

Cost:
- O(n) time, O(k) space where k is the number of characters in the map.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestNoRepeatSubstring(t *testing.T) {
	tests := []struct {
		in       string
		expected int
	}{
		{"", 0},
		{"a", 1},
		{"aa", 1},
		{"ab", 2},
		{"aba", 2},
		{"abb", 2},
		{"abc", 3},
		{"abccde", 3},
		{"aabccbb", 3},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			noRepeatSubstring(tt.in),
		)
	}
}

func noRepeatSubstring(s string) int {
	maxLength, start := 0, 0

	// char remembers the last index of each character we have processed.
	char := map[string]int{}

	for end := range s {
		endChar := string(s[end])

		// if we have seen the character before, update the start index to
		// skip visited characters.
		if _, ok := char[endChar]; ok {
			start = char[endChar] + 1
		}

		// cache the current character' index.
		char[endChar] = end

		// update the maximum length.
		maxLength = common.Max(maxLength, end-start+1)
	}

	return maxLength
}
