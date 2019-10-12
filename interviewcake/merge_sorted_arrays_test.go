// Problem:
// Merge two sorted arrays.
//
// Example:
// Given:  []int{1, 3, 5}, []int{2, 4, 6}
// Return: []int{1, 2, 3, 4, 5, 6}
//
// Solution:
// Use two pointers approach to iterate through the list and keep appending the
// smaller value to a new list.
//
// Cost:
// O(n) time, O(n) space.

package interviewcake

import (
	"reflect"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      []int
		expected []int
	}{
		{
			[]int{},
			[]int{},
			[]int{},
		},
		{
			[]int{},
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 3, 5},
			[]int{2, 4, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1, 3, 5},
			[]int{2, 4, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for _, tt := range tests {
		result := mergeSortedArray(tt.in1, tt.in2)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

func mergeSortedArray(a1, a2 []int) []int {
	out := []int{}

	// keep two "pointers" at index 0 and move up accordingly as one get
	// merged in.
	i, j := 0, 0
	for i < len(a1) && j < len(a2) {
		if a1[i] < a2[j] {
			out = append(out, a1[i])
			i++
		} else {
			out = append(out, a2[j])
			j++
		}
	}

	// if we get here, one array must have bigger size than the other. could
	// figure out which one is it then copy the rest of its to our final one.
	if i < len(a1) {
		out = append(out, a1[i:]...)
	}

	if j < len(a2) {
		out = append(out, a2[j:]...)
	}

	return out
}
