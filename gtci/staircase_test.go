/*
Problem:
- Given a stair with n steps, count how many possible ways to reach the top
  where at each step you can either take 1, 2 or 3 steps.

Example:
- Input: 3
  Output: 4
  Explanation: 4 ways are 1-1-1, 1-2, 2-1, 3
- Input: 4
  Output: 7
  Explanation: 7 ways are 1-1-1-1, 1-1-2, 1-2-1, 2-1-1, 2-2, 1-3, 3-1

Approach:
- Similar to Fibonacci numbers problem, every count is a sum of three
  preceding numbers.

Cost:
- Brute-force: O(n^3) time, O(n) space.
- Top-down: O(n) time, O(n) space.
- Bottom-up: O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestCountSteps(t *testing.T) {
	tests := []struct {
		in       int
		expected int
	}{
		{3, 4},
		{4, 7},
		{5, 13},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			countStepsBF(tt.in),
		)

		common.Equal(
			t,
			tt.expected,
			countStepsTD(tt.in),
		)

		common.Equal(
			t,
			tt.expected,
			countStepsBU(tt.in),
		)
	}
}

func countStepsBF(n int) int {
	// if there is less than 2 steps, there must only be 1 way.
	if n < 2 {
		return 1
	}

	// if there are 2 steps, then there are 2 ways that is 1-1 or 2.
	if n == 2 {
		return 2
	}

	return countStepsBF(n-1) + countStepsBF(n-2) + countStepsBF(n-3)
}

func countStepsTD(n int) int {
	memo := make([]int, n+1)
	return countStepsMemoRecur(memo, n)
}

func countStepsMemoRecur(memo []int, n int) int {
	if n < 2 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = countStepsMemoRecur(memo, n-1) + countStepsMemoRecur(memo, n-2) + countStepsMemoRecur(memo, n-3)
	return memo[n]
}

func countStepsBU(n int) int {
	if n < 2 {
		return 1
	}

	if n == 2 {
		return 2
	}

	n1, n2, n3, tmp := 1, 1, 2, 0
	for i := 3; i < n+1; i++ {
		tmp = n1 + n2 + n3
		n1 = n2
		n2 = n3
		n3 = tmp
	}

	return n3
}
