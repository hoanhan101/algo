/*
Problem:
- Given a list of parentheses, determine if it is valid.

Example:
- Input: []string{"(", ")", "[", "]", "{", "}"}
  Output: true
- Input: []string{"(", "[", ")", "]"}
  Output: false

Approach:
- Use a stack push the opening parenthesis and pop the last inserted one when
  we encounter a closing parenthesis and check if it is a valid match.

Cost:
- O(n) time, O(n) space.
*/

package leetcode

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		in       []string
		expected bool
	}{
		{[]string{}, true},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			isValid(tt.in),
		)
	}
}

func isValid(s []string) bool {
	m := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	stack := newStack()

	for _, p := range s {
		// if p is an opening parenthesis, push it to the stack. otherwise, if
		// is a closing parenthesis, pop the stack and check if they're matching
		if _, ok := m[p]; ok {
			stack.push(p)
		} else if stack.size() == 0 || m[stack.pop()] != p {
			return false
		}
	}

	return stack.size() == 0
}

// stack represent a stack data structure.
type stack struct {
	items []string
}

// newStack return a new Stack.
func newStack() *stack {
	items := []string{}
	return &stack{items: items}
}

// push a new item onto the stack.
func (s *stack) push(v string) {
	s.items = append(s.items, v)
}

// pop remove and return the last item.
func (s *stack) pop() string {
	// get the last item.
	i := len(s.items) - 1
	top := s.items[i]

	// remove the last item.
	s.items = s.items[:i]

	return top
}

// size return the stack's size.
func (s *stack) size() int {
	return len(s.items)
}
