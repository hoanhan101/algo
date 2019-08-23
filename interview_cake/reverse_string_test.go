// Problem:
// Given a list of string, reverse its order in place.
//
// Example:
// Given:  []string{"a", "b", "c", "d"},
// Return: []string{"d", "c", "b", "a"}
//
// Solution:
// Use two pointers approach to swap two values on both ends as we move toward
// the middle.
//
// Cost:
// O(n) time, O(1) space.

package main

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		in       []string
		expected []string
	}{
		{[]string{}, []string{}},
		{[]string{"a"}, []string{"a"}},
		{[]string{"a", "b"}, []string{"b", "a"}},
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}},
		{[]string{"a", "b", "c", "d"}, []string{"d", "c", "b", "a"}},
	}

	for _, tt := range tests {
		result := reverseString(tt.in)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

func reverseString(list []string) []string {
	// check edge cases.
	if len(list) == 0 || len(list) == 1 {
		return list
	}

	// initialize start and end index pointer.
	start := 0
	end := len(list) - 1

	for start < end {
		// swap 2 character using a temp variable.
		tmp := list[start]
		list[start] = list[end]
		list[end] = tmp

		// move the cursor toward the middle.
		start++
		end--
	}

	return list
}
