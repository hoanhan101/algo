// Problem:
// Given a list of integers, return the highest product of three numbers.
//
// Example:
// Given:  []int{-10, -10, 1, 3, 2}
// Return: 300
//
// Solution:
// Use a greedy approach to keep track of the current highest, current lowest,
// highest of three, highest of two, lowest of two.
//
// Cost:
// O(n) time, O(1) space.

package main

import (
	"reflect"
	"testing"
)

func TestHighestProductOfThree(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{-10, -10, 1, 3, 2}, 300},
		{[]int{1, 10, -5, 1, -100}, 5000},
	}

	for _, tt := range tests {
		result := highestProductOfThree(tt.in)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

func highestProductOfThree(list []int) int {
	if len(list) <= 3 {
		return 0
	}

	_, highest := mimax(list[0], list[1])
	lowest, _ := mimax(list[0], list[1])
	highestTwo := list[0] * list[1]
	lowestTwo := list[0] * list[1]
	highestThree := list[0] * list[1] * list[2]

	for i := 2; i < len(list); i++ {
		current := list[i]

		_, highestThree = mimax(highestThree, current*highestTwo, current*lowestTwo)
		_, highestTwo = mimax(highestTwo, current*highest, current*lowest)
		lowestTwo, _ = mimax(lowestTwo, current*highest, current*lowest)
		_, highest = mimax(highest, current)
		lowest, _ = mimax(lowest, current)
	}

	return highestThree
}
