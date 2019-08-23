// Problem:
// Given a sentence (string), return its word count map.
//
// Example:
// Given: "Cliff finished his cake and paid the bill. Bill finished his cake at the edge of the cliff.",
// Return: map[string]int{"cliff": 2, "finished": 2, "his": 2, "cake": 2, "and": 1, "paid": 1, "the": 3, "bill": 2, "at": 1, "edge": 1, "of": 1}
//
// Solution:
// Get rid of special characters using regex then split the sentence by space.
// Use a hashmap to keep counts.
//
// Cost:
// O(n) time, O(n) space.

package main

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestBuildWordCloud(t *testing.T) {
	tests := []struct {
		in       string
		expected map[string]int
	}{
		{
			"Cliff finished his cake and paid the bill. Bill finished his cake at the edge of the cliff.",
			map[string]int{"cliff": 2, "finished": 2, "his": 2, "cake": 2, "and": 1, "paid": 1, "the": 3, "bill": 2, "at": 1, "edge": 1, "of": 1},
		},
	}

	for _, tt := range tests {
		result := buildWordCloud(tt.in)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

func buildWordCloud(in string) map[string]int {
	m := map[string]int{}

	// get rid of all special characters and numbers using regex.
	s := regexp.MustCompile(`[^a-zA-Z]+`).ReplaceAllString(in, " ")

	// split it by space.
	words := strings.Split(s, " ")

	// iterate through the word list and update its count.
	for _, v := range words[:len(words)-1] {
		if _, ok := m[strings.ToLower(v)]; ok {
			m[strings.ToLower(v)]++
		} else {
			m[strings.ToLower(v)] = 1
		}
	}

	return m
}
