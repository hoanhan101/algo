// Problem:
// Given a list of words, return an index of a rotation point.
//
// Solution:
// Use a binary search approach to search for word. The rotation point is
// converged in the middle.
//
// Cost: O(logn) time, O(1) space.

package main

import (
	"reflect"
	"testing"
)

func TestFindRotationPoint(t *testing.T) {
	tests := []struct {
		in       []string
		expected int
	}{
		{
			[]string{
				"ptolemaic",
				"retrograde",
				"supplant",
				"undulate",
				"xenoepist",
				"asymptote", // <-- rotates here!
				"babka",
				"banoffee",
				"engender",
				"karpatka",
				"othellolagkage",
			},
			5,
		},
		{
			[]string{
				"xenoepist",
				"asymptote", // <-- rotates here!
				"babka",
				"banoffee",
				"engender",
				"karpatka",
				"othellolagkage",
			},
			1,
		},
	}

	for _, tt := range tests {
		result := findRotationPoint(tt.in)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

func findRotationPoint(words []string) int {
	firstWord := words[0]
	floorIndex := 0
	ceilingIndex := len(words) - 1

	for floorIndex < ceilingIndex {
		guessIndex := floorIndex + (ceilingIndex-floorIndex)/2

		// if the guess word comes after the first word or the first word, then
		// go right. otherwise, go left.
		if words[guessIndex] >= firstWord {
			floorIndex = guessIndex
		} else {
			ceilingIndex = guessIndex
		}

		// floorIndex and ceilingIndex have converged at rotation point.
		if floorIndex+1 == ceilingIndex {
			return ceilingIndex
		}
	}

	// FIXME - returns 0 if there is no rotation point.
	return 0
}
