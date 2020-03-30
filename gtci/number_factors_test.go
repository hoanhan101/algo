/*
Problem:
- Given a number n, count how many possible ways to calculate n
  as the sum of 1, 3, 4.

Example:
- Input: 4
  Output: 4
  Explanation: 4 ways are 1-1-1-1, 1-3, 3-1, 4
- Input: 5
  Output: 6
  Explanation: 4 ways are 1-1-1-1-1-1, 1-1-3, 1-3-1, 3-1-1, 1-4, 4-1

Approach:
- For every number, we can either subtract 1, 3, or 4 in a recursive way.

Cost:
- Brute-force: O(n^3) time, O(n) space.
- Top-down: O(n) time, O(n) space.
- Bottom-up: O(n) time, O(n) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestCountNumberFactors(t *testing.T) {
	tests := []struct {
		in       int
		expected int
	}{
		{4, 4},
		{5, 6},
		{6, 9},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			countNumberFactorsBF(tt.in),
		)

		common.Equal(
			t,
			tt.expected,
			countNumberFactorsTD(tt.in),
		)

		common.Equal(
			t,
			tt.expected,
			countNumberFactorsBU(tt.in),
		)
	}
}

func countNumberFactorsBF(n int) int {
	// if n is less and equal than 2, there is only 1 way to subtract 1.
	if n <= 2 {
		return 1
	}

	// if n is 3, there are 2 ways as 1-1-1, 3 would work.
	if n == 3 {
		return 2
	}

	return countNumberFactorsBF(n-1) + countNumberFactorsBF(n-3) + countNumberFactorsBF(n-4)
}

func countNumberFactorsTD(n int) int {
	memo := make([]int, n+1)
	return countNumberFactorsMemoRecur(memo, n)
}

func countNumberFactorsMemoRecur(memo []int, n int) int {
	if n <= 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = countNumberFactorsMemoRecur(memo, n-1) + countNumberFactorsMemoRecur(memo, n-3) + countNumberFactorsMemoRecur(memo, n-4)
	return memo[n]
}

func countNumberFactorsBU(n int) int {
	tabu := make([]int, n+1)
	tabu[0] = 1
	tabu[1] = 1
	tabu[2] = 1
	tabu[3] = 2

	for i := 4; i < n+1; i++ {
		tabu[i] = tabu[i-1] + tabu[i-3] + tabu[i-4]
	}

	return tabu[n]
}
