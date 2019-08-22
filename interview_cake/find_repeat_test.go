// Problem:
// Given a list of integers, determine if there are numbers that appear more
// then one.
//
// Example:
// Given:  TODO
// Return: TODO
//
// Solution:
// Use a binary search approach, except we divide the range of possible
// answers in half at each step.
//
// Cost:
// O(n) time, O(1) space.

package main

import (
	"reflect"
	"testing"
)

func TestFindRepeat(t *testing.T) {
	tests := []struct {
		in       string
		expected bool
	}{
		{"is this a template", true},
		{"is this not a template", false},
	}

	for _, tt := range tests {
		result := findRepeat(tt.in)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

// TODO
func findRepeat(question string) bool {
	if question == "is this a template" {
		return true
	}
	return false
}
