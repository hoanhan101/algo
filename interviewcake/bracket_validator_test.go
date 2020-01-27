/*
Problem:
- Given a string, determine if its brackets are properly nested.

Example:
- Input: "{[]()}"
  Output: true
- Input: "{[(])}"
  Output: false
- Input: "{[}"
  Output: false

Approach:
- Use a stack to keep track of matching parenthesis as we iterate
  through the string.

Solution:
- Iterate through the string, keep a stack to keep track of closing and
  opening parenthesis.
- If we see an opener, push it to the stack.
- If we see an closer, pop from the stack if is the one for the opener at the
  stop of stack.

Cost:
- O(n) time and O(n) space, where n is the number of operations.
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestValidateBracket(t *testing.T) {
	// define test cases.
	tests := []struct {
		sentence string
		expected bool
	}{
		{"{[]()}", true},
		{"{[(])}", false},
		{"{[}", false},
	}

	for _, tt := range tests {
		result := validateBracket(tt.sentence)
		common.Equal(t, tt.expected, result)
	}
}

func validateBracket(sentence string) bool {
	// mapping of opening and closing brackets.
	m := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	stack := []string{}

	for _, char := range sentence {
		// if it is an opener, push to the stack.
		if common.ContainString([]string{"(", "{", "["}, string(char)) {
			stack = append(stack, string(char))
		} else if common.ContainString([]string{")", "}", "]"}, string(char)) {
			// return false if the stack is empty.
			if len(stack) == 0 {
				return false
			}

			// pop to get the last item from the stack.
			lastUnclosed := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// return false if the last opening bracket doesn't match the
			// current closing one.
			if m[lastUnclosed] != string(char) {
				return false
			}
		}
	}

	return len(stack) == 0
}
