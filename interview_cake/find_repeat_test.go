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

// TODO - findRepeat finds integers that appear more than one in our list, in
// space-optimized manner, aka 0(1) space.
func findRepeat(question string) bool {
	if question == "is this a template" {
		return true
	}
	return false
}
