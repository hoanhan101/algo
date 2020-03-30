/*
Problem:
- Calculate the nth Fibonacci numbers.

Example:
- Input: 5
  Output: 5
- Input: 6
  Output: 8
- Input: 7
  Output: 13

Top-down approach:
-

Bottom-up approach:
-

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFibonacciNumbers(t *testing.T) {
	tests := []struct {
		in       int
		expected int
	}{
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			bruteForceFib(tt.in),
		)

		common.Equal(
			t,
			tt.expected,
			topDownFib(tt.in),
		)

		common.Equal(
			t,
			tt.expected,
			bottomUpFib(tt.in),
		)
	}
}

func bruteForceFib(n int) int {
	if n < 2 {
		return n
	}

	return bruteForceFib(n-1) + bruteForceFib(n-2)
}

func topDownFib(n int) int {
	// memo caches the calculated values in an array where each index represents
	// the nth number and its index represents the calculated value.
	memo := make([]int, n+1)
	return memoRecur(memo, n)
}

func memoRecur(memo []int, n int) int {
	if n < 2 {
		return n
	}

	// return the value immediately if we've already calculated it
	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = memoRecur(memo, n-1) + memoRecur(memo, n-2)

	return memo[n]
}

func bottomUpFib(n int) int {
	if n < 2 {
		return n
	}

	// initialize n1, n2 as the first 2 numbers in the Fibonacci sequence and
	// use them to cache the last two as we go.
	n1, n2, tmp := 0, 1, 0
	for i := 2; i < n+1; i++ {
		tmp = n1 + n2
		n1 = n2
		n2 = tmp
	}

	return n2
}
