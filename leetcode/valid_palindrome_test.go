// Problem:
// Given a string, determine if it is a palindrome, considering only
// alphanumeric characters and ignoring cases.
//
// Example:
// Input: "A man, a plan, a canal: Panama" // Output: true // Input: "race a car" // Output: false
//
// Solution:
// Use two pointers approach where one points to the start of the string while
// the other points at the end.
// Move them toward each other and compare if they're the same characters while
// skipping on-alphanumeric characters and ignoring cases.
//
// Cost:
// O(n) time, O(1) space.

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
