/*
Problem:
- Given a list of stock prices (integer) in chronological order, return the
  max profit from buying at earlier time and selling at later time.

Example:
- Input: []int{10, 7, 5, 8, 11, 9}
  Output: 6, because one can buy at 5 and sell at 11

Approach:
- Use a greedy approach to keep track of the minimum price and the maximum
  profit for each value while iterating through the list.

Solution:
- Initialize a minimum price to be the first price in the list and a maximum
  profit to be the first possible profit that we could trade.
- Iterate through the list and keep track of:
  - the current price
  - the potential profit: current price - minimum price
  - the maximum profit: the larger profit between the current maximum and the potential one
  - the minimum price: the smaller price between the current minimum and the current price.
- Return the maximum profit we found in the end.

Cost:
- O(n) time, O(1) space.
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestGetMaxProfit(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{10}, 0},
		{[]int{10, 10}, 0},
		{[]int{10, 7, 5, 8, 11, 9}, 6},
		{[]int{10, 2, 5, 4, 7, 1}, 5},
		{[]int{10, 7, 2, 1}, -1},
	}

	for _, tt := range tests {
		result := getMaxProfit(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func getMaxProfit(stocks []int) int {
	// return 0 if the input is invalid
	if len(stocks) <= 2 {
		return 0
	}

	// initialize minPrice to be the first price in the list and maxProfit to
	// be the first possible profit that we could trade.
	minPrice := stocks[0]
	maxProfit := stocks[1] - stocks[0]

	for i := 1; i < len(stocks); i++ {
		currentPrice := stocks[i]
		potentialProfit := currentPrice - minPrice
		maxProfit = common.Max(potentialProfit, maxProfit)
		minPrice = common.Min(minPrice, currentPrice)
	}

	return maxProfit
}
