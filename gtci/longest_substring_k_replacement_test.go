/*
Problem:
- Given a string, if you are allowed to replace  no more than k letters with
  any letter, find the length of the longest substring having the same letters
  after replacement.

Example:
- Input: string="aabccbb", k=2
  Output: 5
  Explanation: Longest substring is "bbbbb" after replacing 2 c with b.
- Input: string="abbcb", k=1
  Output: 4
  Explanation: Longest substring is "bbbb" after replacing 1 c with b.

Approach:
- Use a hashmap to remember the frequency of each character we have seen.
- As we iterate through the string and add character to the window, we
  also keep track of the maximum repeating character count.
- Shrink the window accordingly as we are not allowed to replace more than
  k characters.

Cost:
- O(n) time, O(1) space since there are only 26 characters in the alphabet.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestLongestSubstringKReplacement(t *testing.T) {
	tests := []struct {
		in1      string
		in2      int
		expected int
	}{
		{"aabccbb", 2, 5},
		{"abbcb", 1, 4},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			longestSubstringKReplacement(tt.in1, tt.in2),
		)
	}
}

func longestSubstringKReplacement(s string, k int) int {
	maxLength, start, maxRepeatCharCount := 0, 0, 0

	// char keeps track of characters' frequencies.
	char := map[string]int{}

	for end := range s {
		// insert characters into the frequencies map.
		endChar := string(s[end])
		if _, ok := char[endChar]; !ok {
			char[endChar] = 0
		}
		char[endChar]++

		// update max repeating character count.
		maxRepeatCharCount = common.Max(maxRepeatCharCount, char[endChar])

		// shrink the window as we are not allowed to replace more than k
		// characters.
		if end-start+1-maxRepeatCharCount > k {
			startChar := string(s[start])
			char[startChar]--
			start++
		}

		// update the maximum length at each step.
		maxLength = common.Max(maxLength, end-start+1)
	}

	return maxLength
}
