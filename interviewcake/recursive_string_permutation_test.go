/*
Problem:
- Write a recursive function for generating all permutations of an input
  string. Assume that every character in the string is unique.

Example:
- Input: "cat"
  Output: []set{"cat", "cta", "act", "atc", "tca", "tac"}

Approach:
- Get all the permutations for all characters before the last one.
- Put the last character in all possible position for each of these
  permutations.

Solution:
- Initialize permutations as a set.
- If there is only one character in a string (or less than), return it
  immediately.
- Get the last character and all the characters before it.
- Get all permutations for all characters expect the last one.
- Iterate through the list of permutations before the last character
  and put the last character in all possible position for each of these
  permutations.

Cost:
- O(n) time, O(n) space.
*/

package interviewcake

import (
	"strings"
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestPermuteString(t *testing.T) {
	tests := []struct {
		in       string
		expected map[string]bool
	}{
		{"c", map[string]bool{"c": true}},
		{"ca", map[string]bool{"ca": true, "ac": true}},
		{"cat", map[string]bool{"cat": true, "cta": true, "act": true, "atc": true, "tca": true, "tac": true}},
		{"caa", map[string]bool{"caa": true, "aca": true, "aac": true}},
	}

	for _, tt := range tests {
		result := permuteString(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func permuteString(in string) map[string]bool {
	// initialize permutations as a set.
	permutations := make(map[string]bool)

	// if there is only one character in a string (or less than), return it
	// immediately.
	if len(in) <= 1 {
		permutations[in] = true
		return permutations
	}

	// get the last character and all the characters before it.
	lastChar := string(in[len(in)-1])
	beforeLastChar := string(in[:len(in)-1])

	// get all permutations for all characters expect the last one.
	permutationsBeforeLastChar := permuteString(beforeLastChar)

	// put the last character in all possible position for each of these
	// permutations.
	for permutation := range permutationsBeforeLastChar {
		for position := range beforeLastChar {
			// by keeping track of the position, we can put the last character
			// in between the permutation.
			s := []string{permutation[:position], lastChar, permutation[position:]}
			p := strings.Join(s, "")

			// put in the set.
			permutations[p] = true

			// at position 0, permutation[:0] is an empty string. hence,
			// just joining the permutation with the last character is
			// enough.
			if position == 0 {
				s := []string{permutation, lastChar}
				p := strings.Join(s, "")
				permutations[p] = true
			}
		}
	}

	return permutations
}
