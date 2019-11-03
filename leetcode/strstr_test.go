/*
Problem:
- Implement strstr() that finds the first occurrence of the substring
  needle in the string haystack. It returns -1 if needle is not part of the
  haystack.

Example:
- Input: haystack = "aaabacd", needle = "ba"
  Output: 3, because needle "ba" starts at index 3 in the haystack.

Approach:
- Scan the needle with the haystack from its first position and start matching
  all subsequent letters one by one.
- If one letter does not match, start again with the next position in the
  haystack.

Cost:
- O(nm) time, O(1) space, where n is the length of haystack while m is the
  length of needle.
*/

package leetcode

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestStrstr(t *testing.T) {
	tests := []struct {
		in1      string
		in2      string
		expected int
	}{
		{"", "", -1},
		{"", "a", -1},
		{"dani", "a", 1},
		{"danidani", "a", 1},
		{"xxxdani", "dani", 3},
		{"xxxdanixxx", "dani", 3},
	}

	for _, tt := range tests {
		result := strstr(tt.in1, tt.in2)
		common.Equal(t, tt.expected, result)
	}
}

func strstr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		for j := 0; j < len(haystack); j++ {
			// if the needle matches the haystack, meaning the index pointer
			// reaches the end of the needle, return the index pointer.
			if j == len(needle) {
				return i
			}

			// return -1 if needle's length is greater than haystack's.
			if i+j == len(haystack) {
				return -1
			}

			// start again if one of the letter in the needles does not match
			// one in the haystack.
			if []rune(needle)[j] != []rune(haystack)[i+j] {
				break
			}
		}
	}

	return -1
}
