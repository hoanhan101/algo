// Problem:
// Given a string, check if its permutation is a palindrome.
//
// Example:
// Given: "ivicc" or "civic"
// Return: true
//
// Solution:
// To determine if a permutation is a palindrome, need to check if each
// character in the string appears an even number of times, allowing for
// only one character to appear an odd time, that is the middle one.
// Use a hashmap to add a character if we haven't seen it and remove it
// if it's already there.
// After the iteration, if we're left with less or equal than a character in
// the map, we have a palindrome.
//
// Cost:
// O(n) time, O(1) space.
// The space complexity is O(n) due to the hashmap, but since there are
// only a constant number of characters in Unicode, we could treat it
// as O(1).

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
		{"", true},
		{"c", true},
		{"cc", true},
		{"ca", false},
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

	// delete if we've seen it, else add it with a count 1.
	for _, v := range in {
		if _, ok := m[string(v)]; ok {
			delete(m, string(v))
		} else {
			m[string(v)] = 1
		}
	}

	// if we're left with less or equal than a pair, we have a palindrome.
	if len(m) <= 1 {
		return true
	}

	return false
}
