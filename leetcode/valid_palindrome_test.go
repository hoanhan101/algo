/*
Problem:
- Given a string, determine if it is a palindrome, considering only
  alphanumeric characters and ignoring cases.

Example:
- Input: "A man, a plan, a canal: Panama"
  Output: true
- Input: "race a car"
  Output: false

Approach:
- Use two pointers approach that have one point to the start of the string and
  the other point at the end.
- Move them toward each other and compare if they're the same characters while
  skipping non-alphanumeric characters and ignoring cases.

Solution:
- Have a pointer point to the start of the string and the other point at the end.
- Make the string lower case.
- While the start one is less the end one, ignore non-alphanumeric characters
  and move them toward each other.

Cost:
- O(n) time, O(1) space.
*/

package leetcode

import (
	"strings"
	"testing"
	"unicode"

	"github.com/hoanhan101/algo/common"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		in       string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
	}

	for _, tt := range tests {
		result := isPalindrome(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1

	// make the input string lower case.
	s = strings.ToLower(s)

	for start < end {
		// if the letter at the start and end pointers are non-alphanumeric
		// then skip them.
		for !unicode.IsLetter([]rune(s)[start]) {
			start++
		}
		for !unicode.IsLetter([]rune(s)[end]) {
			end--
		}

		// return false immediately if two corresponding characters are not
		// matching.
		if []rune(s)[start] != []rune(s)[end] {
			return false
		}

		start++
		end--
	}

	return true
}
