/*
Problem:
- Given a sentence as string, and the position of an opening parenthesis
  position, find the matching closing one position.

Example:
- Input: "I ((like) (nesting) parenthesis)", opening parenthesis position = 2
  Output: 31, because the matching parenthesis of the one in position 2 is at
  index 31.

Approach:
- Iterate through the string and keep a count of matching parenthesis at each
  step.

Solution:
- Return -1 if the string is empty or the opening parenthesis position is
  larger than the string size.
- Iterate through the string, starting at the opening parenthesis position and
  keep a count of it.
- Increase the count as we meet the opening parenthesis and decrement it as we
  meet the closing parenthesis.
- When the count if 0, it is the matching end.

Cost:
- O(n) time, O(1) space.
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestMatchParenthesis(t *testing.T) {
	// define test cases..
	tests := []struct {
		sentence   string
		startIndex int
		expected   int
	}{
		{"", 0, -1},
		{"", 1, -1},
		{"(", 0, -1},
		{"()", 0, 1},
		{"())", 0, 1},
		{"(()", 0, -1},
		{"I (like (nesting) parenthesis)", 8, 16},
		{"I ((like) (nesting) parenthesis)", 2, 31},
	}

	for _, tt := range tests {
		common.Equal(t, tt.expected, matchParenthesis(tt.sentence, tt.startIndex))
	}
}

func matchParenthesis(sentence string, startIndex int) int {
	// return -1 if the string is empty or the opening parenthesis position is
	// larger than the string size.
	if len(sentence) == 0 || startIndex > len(sentence) {
		return -1
	}

	count := 0
	for i := startIndex + 1; i < len(sentence); i++ {
		char := sentence[i]

		if string(char) == "(" {
			count++
		} else if string(char) == ")" {
			if count == 0 {
				return i
			}

			count--
		}
	}

	// return -1 if there is no matching parenthesis.
	return -1
}
