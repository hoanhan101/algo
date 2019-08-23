// Problem:
// Given an unsorted list scores (integer) and a highest possible score
// (integer), return a sorted list utilizing that fact.
//
// Example:
// Given:  []int{37, 89, 41, 65, 91, 53}, 100
// Return: []int{91, 89, 65, 53, 41, 37}
//
// Solution:
// Build a list of scores where their indices represent the scores itself and
// their values represent how many time the score appear in the list.
// Utilize the highest score to allocate that fix-sized list ahead of time.
// Iterate backward through that list and add them back in order.
//
// Cost: // O(n) time, O(n) space.

package main

import (
	"reflect"
	"testing"
)

func TestSortScores(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      int
		expected []int
	}{
		{[]int{}, 0, []int{}},
		{[]int{37, 89, 41, 65, 91, 53}, 100, []int{91, 89, 65, 53, 41, 37}},
		{[]int{37, 89, 41, 65, 91, 53, 65}, 100, []int{91, 89, 65, 65, 53, 41, 37}},
	}

	for _, tt := range tests {
		result := sortScores(tt.in1, tt.in2)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

func sortScores(scores []int, highestPossibleScore int) []int {
	// utilize highest score to create a fix-sized list of score counts.
	counts := make([]int, highestPossibleScore+1)
	out := make([]int, 0)

	// populate the scores where indices represent scores and values represent
	// how many time the score appears.
	for _, score := range scores {
		counts[score]++
	}

	// iterate backward and add the score to the output list so that it is
	// sorted from highest to lowest.
	for i := len(counts) - 1; i > -1; i-- {
		count := counts[i]

		for j := 0; j < count; j++ {
			out = append(out, i)
		}
	}

	return out
}
