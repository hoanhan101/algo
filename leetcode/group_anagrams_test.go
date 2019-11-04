/*
Problem:
- Given an array of strings, group anagrams together.
- All inputs will be in lowercase.
- The order of your output does not matter.

Example:
- Input: []string{"eat", "tea", "tan", "ate", "nat", "bat"}
  Output: [][]string{
	  []string{"ate", "eat", "tea"},
	  []string{"nat","tan"},
	  []string{"bat"},
  }

Approach:
- Two strings are anagrams if and only if their character counts
  (respective number of occurrences of each character) are the same.

Cost:
- O(n) time, O(n) space.
*/

package leetcode

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

// func TestGroupAnagrams(t *testing.T) {
// 	tests := []struct {
// 		in       []string
// 		expected [][]string
// 	}{
// 		{
// 			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
// 			[][]string{
// 				[]string{"ate", "eat", "tea"},
// 				[]string{"nat", "tan"},
// 				[]string{"bat"},
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		result := groupAnagrams(tt.in)
// 		common.Equal(t, tt.expected, result)
// 	}
// }

func TestCountCharacters(t *testing.T) {
	tests := []struct {
		in       string
		expected map[string]int
	}{
		{"", map[string]int{}},
		{"ab", map[string]int{"a": 1, "b": 1}},
		{"aba", map[string]int{"a": 2, "b": 1}},
		{"ababb", map[string]int{"a": 2, "b": 3}},
	}

	for _, tt := range tests {
		result := countCharacters(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

// countCharacters counts the number of occurrences of each character for a
// given word.
func countCharacters(word string) map[string]int {
	m := map[string]int{}

	for _, c := range word {
		if _, ok := m[string(c)]; ok {
			m[string(c)]++
		} else {
			m[string(c)] = 1
		}
	}

	return m
}

// FIXME: go doesn't allow to have a map of map[string]int to []string.
// func groupAnagrams(anagrams []string) [][]string {
// 	m := map[map[string]int][]string{}

// 	for _, word := range anagrams {
// 		m[countCharacters(word)] = append(m[countCharacters(word)], word)
// 	}

// 	out := []string{}
// 	for _, v := range m {
// 		out := append(out, v)
// 	}

// 	return out
// }
