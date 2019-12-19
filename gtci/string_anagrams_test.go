/*
Problem:
- Given a string and a pattern, find all anagrams of the pattern in the given string
- Return them as a list of anagrams' starting indices.

Example:
- Input: string="ppqp", pattern="pq"
  Output: []int{1, 2}
- Input: string="abbcabc", pattern="abc"
  Output: []int{2, 3, 4}

Approach:
- Similar to permutation in string problem, except we will store the starting indices
  of the anagrams of the pattern.

Cost:
- O(n) time, O(m) space where m is the number of distinct characters in the map.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFindStringAnagrams(t *testing.T) {
	tests := []struct {
		in1      string
		in2      string
		expected []int
	}{
		{"ppqp", "pq", []int{1, 2}},
		{"abbcabc", "abc", []int{2, 3, 4}},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			findStringAnagrams(tt.in1, tt.in2),
		)
	}
}

func findStringAnagrams(s string, pattern string) []int {
	out := []int{}
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
			out = append(out, start)
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

	return out
}
