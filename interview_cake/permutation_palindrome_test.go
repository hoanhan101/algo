package main

import (
	"reflect"
	"testing"
)

func TestHasPalindromePermutation(t *testing.T) {
	tests := []struct {
		in       string
		expected bool
	}{
		{"civic", true},
		{"ivicc", true},
		{"civil", false},
		{"livci", false},
	}

	for _, tt := range tests {
		result := hasPalindromePermutation(tt.in)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

func hasPalindromePermutation(in string) bool {
	m := map[string]int{}

	// delete if we've seen it, else add it
	for _, v := range in {
		if _, ok := m[string(v)]; ok {
			delete(m, string(v))
		} else {
			m[string(v)] = 1
		}
	}

	// if we're left with less or equal than 1 pair, we have a palindrome
	if len(m) <= 1 {
		return true
	}

	return false
}
