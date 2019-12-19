/*
Problem:
- Given a string and a pattern, find out if the string contains any permutation
  of the pattern.

Example:
- Input: string="oidbcaf", pattern="abc"
  Output: true
  Explanation: The string contains "bca" which is a permutation of the given pattern.
- Input: string="odicf", pattern="dc"
  Output: false

Approach:
- Use a hashmap to calculate the frequencies of all characters in the patterns.
- Iterate through the string and add characters in the sliding window.
- At each step,
  - check if we got a complete match
  - shrink the window so that its size is equal to the length of the pattern.

Cost:
- O(n) time, O(m) space where m is the number of distinct characters in the map.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFindPermutation(t *testing.T) {
	tests := []struct {
		in1      string
		in2      string
		expected bool
	}{
		{"oidbcaf", "abc", true},
		{"oidbafc", "abc", false},
		{"odicf", "dc", false},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			findPermutation(tt.in1, tt.in2),
		)
	}
}

func findPermutation(s string, pattern string) bool {
	matched, start := 0, 0

	// char keeps track of characters' frequencies.
	char := map[string]int{}

	// calculate the frequencies of all characters in the pattern.
	for _, c := range pattern {
		if _, ok := char[string(c)]; !ok {
			char[string(c)] = 0
		}

		char[string(c)]++
	}

	for end := range s {
		// if the character matches one in the map, decrease its frequency.
		// if its frequency becomes 0, we have a complete match.
		endChar := string(s[end])
		if _, ok := char[endChar]; ok {
			char[endChar]--

			if char[endChar] == 0 {
				matched++
			}
		}

		// return true immediately if we have all complete matches.
		if matched == len(char) {
			return true
		}

		// shrink the window by one character at a time so that its size is
		// equal to the length of the pattern.
		if end >= len(pattern)-1 {
			startChar := string(s[start])
			start++

			// if the character going out is part of the pattern, put it back in.
			if _, ok := char[startChar]; ok {
				if char[startChar] == 0 {
					matched--
				}

				char[startChar]++
			}
		}

	}

	return false
}
