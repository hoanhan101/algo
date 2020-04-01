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

Cost:
- Brute-force: O(n^2) time, O(n) space.
- Top-down: O(n) time, O(n) space.
- Bottom-up: O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestCalculateFibonacci(t *testing.T) {
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
		{7, 13},
		{8, 21},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			calculateFibBF(tt.in),
		)

		common.Equal(
			t,
			tt.expected,
			calculateFibTD(tt.in),
		)

		common.Equal(
			t,
			tt.expected,
			calculateFibBU(tt.in),
		)
	}
}

func calculateFibBF(n int) int {
	if n < 2 {
		return n
	}

	return calculateFibBF(n-1) + calculateFibBF(n-2)
}

func calculateFibTD(n int) int {
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

func calculateFibBU(n int) int {
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
