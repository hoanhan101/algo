/*
Problem:
- Given an amount of money and a list of coin denominations, compute the
  number of ways to make that amount with these available coins.

Example:
- Input: amount = 4, denominators = []int{1, 2, 3}
  Output: 4, because there are 4 ways to calculate 4 as such:
  4 = 1 + 1 + 1 + 1
  = 1 + 2 + 2
  = 1 + 3
  = 2 + 2

Approach:
- Use a bottom-up approach to build a table of ways for calculating the amount
  using our denominations: the index is the amount, the value at each index
  is the number of ways to get that amount.
- The number of new ways we can make a higher amount when we account for a new
  coin is ways of calculating the difference between the higher amount and
  current coin.

Solution:
- Make a list of ways of calculating the amount, where the index is the
  amount, the value at each index is the number of ways to get that amount.
- Start with the base case that there is one way to calculate 0.
- The number of new ways we can make a higher amount when we account for a
  new coin is ways[higher amount - coin], where the reminder is already
  calculated as we are going bottom-up.

Cost:
- O(n*m) time, O(n) space, where n is the amount of money, m is the number of
  potential denominations.
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestMakeChange(t *testing.T) {
	tests := []struct {
		in1       int
		in2       []int
		expected1 []int
		expected2 int
	}{
		{
			0, []int{},
			[]int{
				1, // 1 way to calculate 0 - base case
			},
			1,
		},
		{
			4, []int{1, 2, 3},
			[]int{
				1, // 1 way to calculate 0
				1, // 1 way to calculate 1
				2, // 2 ways to calculate 2
				3, // 3 ways to calculate 3
				4, // 4 ways to calculate 4
			},
			4,
		},
		{
			5, []int{1, 2},
			[]int{
				1, // 1 way to calculate 0
				1, // 1 way to calculate 1
				2, // 2 way to calculate 2
				2, // 2 ways to calculate 3
				3, // 3 ways to calculate 4
				3, // 3 ways to calculate 5
			},
			3,
		},
		{
			5, []int{1, 3, 5},
			[]int{
				1, // 1 way to calculate 0
				1, // 1 way to calculate 1
				1, // 1 way to calculate 2
				2, // 2 ways to calculate 3
				2, // 2 ways to calculate 4
				3, // 3 ways to calculate 5
			},
			3,
		},
	}

	for _, tt := range tests {
		r1, r2 := makeChange(tt.in1, tt.in2)
		common.Equal(t, tt.expected1, r1)
		common.Equal(t, tt.expected2, r2)
	}
}

func makeChange(amount int, denominations []int) ([]int, int) {
	// make a list of ways of calculating the amount, where the index is the
	// amount, the value at each index is the number of ways to get that
	// amount.
	ways := make([]int, amount+1)

	// start with the base case that there is one way to calculate 0.
	ways[0] = 1

	// the number of new ways we can make a higher amount when we account for a
	// new coin is ways[higherAmount - coin], where the reminder is already
	// calculated as we are going bottom-up.
	for _, coin := range denominations {
		for higherAmount := coin; higherAmount < amount+1; higherAmount++ {
			remainder := higherAmount - coin
			ways[higherAmount] += ways[remainder]
		}
	}

	// NOTE - also returning the calculated table makes it easier to walk-through
	// the logic and test cases by oneself.
	return ways, ways[amount]
}
