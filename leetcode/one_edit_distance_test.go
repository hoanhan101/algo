/*
Problem:
- Given two strings, determine if they are both one edit distance apart.

Example:
- Input: "abcde", "abXde"
  Output: true
  Explanation: Only "c" in S is replaced by "X" in T.
- Input: "abcde", "abcXde"
  Output: true
  Explanation: "X" is inserted between "c" and "d" in S to get T.

Approach:
- Use two-pointer approach to traverse both strings at the same time and
  keep track of count of difference characters.

Solution:
- If the difference between lengths of two strings is more than 1, return
  false immediately because they are not at one distance anyway.
- Iterate through both string at the same time and move both pointers up
  their corresponding characters match as follow:
  - If the length of one string is more than the other, a possible edit is
    to remove a character. We move the pointer in the larger string up.
  - If they are the same length, a possible edit is to change a character.
    In that case, simply move both of them up.
  - Remember to increment the edit count.
- If the character is extra in any string, also increment the edit count.

Cost:
- O(n) time, O(1) space
*/

package leetcode

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestIsOneDistanceApart(t *testing.T) {
	tests := []struct {
		in1      string
		in2      string
		expected bool
	}{

		{"abcde", "abXde", true},
		{"abcde", "abcXde", true},
		{"abcde", "abcdeX", true},
	}

	for _, tt := range tests {
		result := isOneDistanceApart(tt.in1, tt.in2)
		common.Equal(t, tt.expected, result)
	}
}

func isOneDistanceApart(s1, s2 string) bool {
	m := len(s1)
	n := len(s2)

	// if the difference between lengths is more than 1, they can't be at one
	// distance.
	if common.IsMoreThan1Apart(m, n) {
		return false
	}

	count := 0
	i := 0
	j := 0

	for (i < m) && (j < n) {
		// move both pointers up if their corresponding characters match.
		// otherwise, move one up accordingly.
		if s1[i] != s2[j] {
			if count == 1 {
				return false
			}

			// if the length of one string is more than the other, a possible
			// edit is to remove a character. hence, move the pointer in the
			// larger string up.
			// if they are the same length, a possible edit is to change a
			// character. in that case, simply move both of them up.
			if m > n {
				i++
			} else if m < n {
				j++
			} else {

				i++
				j++
			}

			count++
		}

		i++
		j++
	}

	// if the character is extra in any string, also increment the edit count.
	if i < m || j < n {
		count++
	}

	if count == 1 {
		return true
	}

	return false
}
