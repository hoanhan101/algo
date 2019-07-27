package main

import (
	"reflect"
	"testing"
)

func TestTemplate(t *testing.T) {
	tests := []struct {
		in       string
		expected bool
	}{
		{"is this a template", true},
		{"is this not a template", false},
	}

	for _, tt := range tests {
		result := isTemplate(tt.in)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

// isTemplate is just a template function.
func isTemplate(question string) bool {
	if question == "is this a template" {
		return true
	}
	return false
}
