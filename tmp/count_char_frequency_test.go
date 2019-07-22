package main

import (
	"reflect"
	"testing"
)

func TestCountCharFrequency(t *testing.T) {
	tests := []struct {
		in       string
		expected map[string]int
	}{
		{"aabbcc", map[string]int{"a": 2, "b": 2, "c": 2}},
		{"abbc", map[string]int{"a": 1, "b": 2, "c": 1}},
		{"aaaaaa", map[string]int{"a": 6}},
		{"", map[string]int{}},
	}

	for _, tt := range tests {
		result := countCharFrequency(tt.in)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

func countCharFrequency(in string) map[string]int {
	out := map[string]int{}

	for _, v := range in {
		if _, ok := out[string(v)]; ok {
			out[string(v)] += 1
		} else {
			out[string(v)] = 1
		}
	}

	return out
}
