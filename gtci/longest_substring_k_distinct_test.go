/*
Problem:
- Given a string, find the length of the longest substring in it with no more
  than k distinct characters.

Example:
- Input: string="araaci", k=1
  Output: 2
  Explanation: Longest substring with no more than 1 distinct characters is "aa".
- Input: string="araaci", k=2
  Output: 4
  Explanation: Longest substring with no more than 2 distinct characters is "araa".
- Input: string="araaci", k=3
  Output: 5
  Explanation: Longest substring with no more than 3 distinct characters is "araac".

Approach:
- Use a hashmap to remember the frequency of each character we have seen.
- Insert characters until we have k distinct characters in the map to be consider a
  window.
- Shrink the window until there is no more k distinct characters in the map and keep
  updating the maximum window length at each step.

Cost:
- O(n) time, O(k) space where k is the number of characters in the map.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestLongestSubstringKDistinct(t *testing.T) {
	tests := []struct {
		in1      string
		in2      int
		expected int
	}{
		{"", 0, 0},
		{"", 1, 0},
		{"a", 0, 0},
		{"a", 1, 1},
		{"aa", 1, 2},
		{"aa", 2, 2},
		{"ab", 1, 1},
		{"ab", 2, 2},
		{"araaci", 1, 2},
		{"araaci", 2, 4},
		{"araaci", 3, 5},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			longestSubstringKDistinct(tt.in1, tt.in2),
		)
	}
}

func longestSubstringKDistinct(s string, k int) int {
	maxLength, start := 0, 0

	// char keeps track of characters' frequencies.
	char := map[string]int{}

	for end := range s {
		// insert characters until we have k distinct characters.
		endChar := string(s[end])
		if _, ok := char[endChar]; !ok {
			char[endChar] = 0
		}
		char[endChar]++

		// shrink the window until there is no more than k distinct characters.
		for len(char) > k {
			startChar := string(s[start])

			// decrement the frequency of the one going out of the window and
			// remove if its frequency is zero.
			char[startChar]--
			if char[startChar] == 0 {
				delete(char, startChar)
			}

			// increase the start index to move the window ahead by one element.
			start++
		}

		// update the maximum length at each step.
		maxLength = common.Max(maxLength, end-start+1)
	}

	return maxLength
}
