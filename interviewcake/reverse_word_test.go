// Problem:
// Given a list of string that is made up of word but in reverse, return the
// correct order.
//
// Example:
// Given:  []string{"w", "o", "r", "l", "d", "", "h", "e", "l", "l", "o", "", "s", "a", "y"}
// Return: []string{"s", "a", "y", "", "h", "e", "l", "l", "o", "", "w", "o", "r", "l", "d"}
//
// Solution:
// The first pass, we reverse all characters so that we end up with a list of
// words in the right order, but not its characters.
// The second pass, we reverse the order of its characters.
//
// Cost:
// O(n) time, O(1) space.

package interviewcake

import (
	"reflect"
	"testing"
)

func TestReverseWord(t *testing.T) {
	tests := []struct {
		in       []string
		expected []string
	}{
		{
			[]string{"w", "o", "r", "l", "d", "", "h", "e", "l", "l", "o", "", "s", "a", "y"},
			[]string{"s", "a", "y", "", "h", "e", "l", "l", "o", "", "w", "o", "r", "l", "d"},
		},
	}

	for _, tt := range tests {
		result := reverseWord(tt.in)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

func reverseWord(list []string) []string {
	// by reversing all character in the list, we end up with a list of words
	// in the right order but not its characters.
	reverseChar(list, 0, len(list)-1)

	// start keeps track of the start index for each word. it starts with 0 but
	// then gets updated once we find the empty string. then reverse the words
	// characters.
	start := 0
	for i := range list {
		if i == len(list)-1 {
			reverseChar(list, start, i)
		}

		if list[i] == "" {
			reverseChar(list, start, i-1)
			start = i + 1
		}
	}

	return list
}

// reverseChar reverses the list of character for a given start and end index.
func reverseChar(list []string, start int, end int) {
	for start < end {
		tmp := list[start]
		list[start] = list[end]
		list[end] = tmp

		start++
		end--
	}
}
