/*
Problem:
- You are climbing a staircase. It takes n steps to reach to the top.
- Each time you can either climb 1 or 2 steps. In how many distinct ways
  can you climb to the top.

Example:
- Input: 2
  Output: 2
  Explanation: Two ways to climb the top are 1+1 and 2.
- Input: 3
  Output: 3
  Explanation: Two ways to climb the top are 1+1+1, 1+2, and 2+1.

Approach:
- To reach n step, one either advance one step from the n-1 step or
  2 steps from the n-2 step.
- This is basically the same recurrence formula defined by the Fibonacci
  sequence.

Solution:
- Initialize two variables to store 2 previous value.
- Iterate from 2 to n and update these two values at each step.

Cost:
- O(n) time, O(1) space.
*/

package leetcode

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		in       int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			climbStairs(tt.in),
		)
	}
}

func climbStairs(n int) int {
	p, q := 1, 1

	for i := 2; i <= n; i++ {
		tmp := q
		q += p
		p = tmp
	}

	return q
}
