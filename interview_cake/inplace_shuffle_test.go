// Problem:
// Given a list of integers, shuffle its location in-place.
//
// Solution:
// Iterate through the list, swap current value with a value in a randomized
// index that is between the current and last index.
//
// Cost: O(n) time, O(1) space.

package main

import (
	"reflect"
	"testing"
)

func TestInplaceShuffle(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{2, 1, 4, 3, 5}},
	}

	for _, tt := range tests {
		result := inplaceShuffle(tt.in)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

func inplaceShuffle(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	lastIndex := len(list) - 1
	for i := 0; i < len(list); i++ {
		randomIndex := random(i, lastIndex)

		if i != randomIndex {
			tmp := list[i]
			list[i] = list[randomIndex]
			list[randomIndex] = tmp
		}
	}

	return list
}
