/*
Problem:
- Given a string, find the length of the longest substring without repeating characters.

Example:
- Input: "abcabcbb"
  Output: 3
  Explanation: The longest substring is "abc" with the length of 3.
- Input: "bbbbb"
  Output: 1
  Explanation: The longest substring is "b" with the length of 1.

Approach:
- Iterate through the string and keep track of the maximum length of non-repeating
  characters using a hashmap that maps characters to their indices.
- Could skip characters immediately if we found a repeating character.

Solution:
- Initialize a map that maps characters to their indices.
- Initialize a start index and end index to keep track of the start and end of
  a substring.
- Iterate through the string and check if we have seen the current character
  before in the map.
- If so, update the start index.
- Otherwise, cache the current index and update the maximum length if we found
  a larger one.
- Return the maximum length in the end.

Cost:
- O(n) time, O(m) cost where m < n and  n is the length of the string.
*/

package leetcode

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestCalculateLongestSubstring(t *testing.T) {
	tests := []struct {
		in       string
		expected int
	}{
		{"", 0},
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"danixxxdaniiii", 5},
	}

	for _, tt := range tests {
		result := calculateLongestSubstring(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func calculateLongestSubstring(s string) int {
	// m maps characters to their indices.
	m := map[string]int{}

	start := 0
	maxLength := 0

	for end := 0; end < len(s); end++ {
		// if we have seen the character before, update the start index to
		// skip visited characters.
		if m[string(s[end])] >= start {
			start = m[string(s[end])] + 1
		}

		// cache the current character's index at every step.
		m[string(s[end])] = end

		// similar to greedy approach, update max length at each step.
		maxLength = common.Max(end-start+1, maxLength)
	}

	return maxLength
}
