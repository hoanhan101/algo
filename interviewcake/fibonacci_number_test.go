// Problem:
// Write a function that takes an integer n and return the nth Fibonacci
// number. Assume that n is a positive integer.
//
// Example:
// Input: 4
// Output: 3, because the 4th Fibonacci of the sequence 0, 1, 1, 2, 3 is 3
//
// Solution:
// Instead of using a recursive approach, use a bottom-up approach and
// iteratively compute subsequent numbers until we get to n.
//
// Cost:
// O(n) time, O(1) space.

package interviewcake

import (
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	tests := []struct {
		in       int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
	}

	for _, tt := range tests {
		result := fib(tt.in)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

func fib(n int) int {
	// current holds the nth Fibonacci number.
	current := 0

	// return 0 1 or for base cases.
	if n == 0 || n == 1 {
		return n
	}

	// keep track of 2 previous numbers at each step to build the Fibonacci
	// series from the bottom up
	f0 := 0
	f1 := 1

	// to get nth Fibonacci, do n-1 iterations.
	for i := 0; i < n-1; i++ {
		current = f0 + f1
		f0 = f1
		f1 = current
	}

	return current
}
