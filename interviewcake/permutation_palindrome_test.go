/*
Problem:
- Given a string, check if its permutation is a palindrome.

Example:
- Input: "ivicc"
  Output: true
- Input: "civic"
  Output: true

Approach:
- To determine if a permutation is a palindrome, need to check if each
  character in the string appears an even number of times, allowing for
  only one character to appear an odd time, that is the middle one.
- Could use a hashmap store the characters and their number of occurrences.

Solution:
- As we iterate through the string, use a hashmap to add a character if
  we haven't seen it and remove it if it's already there.
- After the iteration, if we're left with less or equal than a character in
  the map, we have a palindrome.

Cost:
- O(n) time, O(1) space.
- The space complexity is O(n) due to the hashmap, but since there are
  only a constant number of characters in Unicode, we could treat it
  as O(1).
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestHasPalindromePermutation(t *testing.T) {
	tests := []struct {
		in       string
		expected bool
	}{
		{"", true},
		{"c", true},
		{"cc", true},
		{"ca", false},
		{"civic", true},
		{"ivicc", true},
		{"civil", false},
		{"livci", false},
	}

	for _, tt := range tests {
		result := hasPalindromePermutation(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func hasPalindromePermutation(in string) bool {
	m := map[string]int{}

	// delete if we've seen it, else add it with a count 1.
	for _, v := range in {
		if _, ok := m[string(v)]; ok {
			delete(m, string(v))
		} else {
			m[string(v)] = 1
		}
	}

	// if we're left with less or equal than a pair, we have a palindrome.
	return len(m) <= 1
}
