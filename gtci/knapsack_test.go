/*
Problem:
- Given a set of items, each with a weight and a value, determine the number of
  each item to include in a collection so that the total weight is less or equal
  than a given limit and the total value is as large as possible.

Example:
- Input: a set of fruit items, with their weights and profits as follow
    fruit  : apple | orange | banana | melon
	weight :   2   |    3   |    1   |   4
	profit :   4   |    5   |    3   |   7
    & knapsack capacity = 5
  Output: banana & melon
  Explanation: banana and melon gives the maximum profit of 10 and weight exactly 5
- Input: weight :   1   |    6   |    10   |   16
	profit :   1   |    2   |    3    |   5
    & knapsack capacity = 7
  Output: 22
  Explanation: 16+6 gives the largest profit and weights exactly 7

Brute-force approach:
- First, calculate the profit for the item at the current index.
- If the total weight does not exceed the capacity, recursively process
  the remaining capacity and items.
- Second, recursively process after excluding the item at the current index.
- Return the higher profit between these two.

Cost:
- Brute-force: O(2^n) time, O(n) space.
- Top-down: O(n*c) time, O(n*c) space where n is the number of items, c is the knapsack capacity.
- Bottom-up: O(n*c) time, O(n*c) space where n is the number of items, c is the knapsack capacity.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestKnapsack(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      []int
		in3      int
		expected int
	}{
		{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, -1, 0},
		{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 0, 0},
		{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 1, 1},
		{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 2, 6},
		{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 3, 10},
		{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 4, 11},
		{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 5, 16},
		{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 6, 17},
		{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 7, 22},
		{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 8, 26},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			knapsackBF(tt.in1, tt.in2, tt.in3),
		)

		common.Equal(
			t,
			tt.expected,
			knapsackTD(tt.in1, tt.in2, tt.in3),
		)

		common.Equal(
			t,
			tt.expected,
			knapsackBU(tt.in1, tt.in2, tt.in3),
		)
	}
}

func knapsackBF(profits, weights []int, capacity int) int {
	return knapsackBFRecur(profits, weights, capacity, 0)
}

func knapsackBFRecur(profits, weights []int, capacity, currentIndex int) int {
	if capacity <= 0 || currentIndex >= len(profits) {
		return 0
	}

	// first, calculate the profit for the item at the current index.
	// if the total weight does not exceed the capacity, recursively process
	// the remaining capacity and items.
	profit1 := 0
	if weights[currentIndex] <= capacity {
		profit1 = profits[currentIndex] + knapsackBFRecur(profits, weights, capacity-weights[currentIndex], currentIndex+1)
	}

	// second, recursively process after excluding the item at the current index.
	profit2 := knapsackBFRecur(profits, weights, capacity, currentIndex+1)

	return common.Max(profit1, profit2)
}

func knapsackTD(profits, weights []int, capacity int) int {
	// since for each recursive call, only capacity and current index change,
	// can have a 2D array for memoization.
	memo := make([][]int, len(profits))
	for i := range memo {
		memo[i] = make([]int, capacity+1)
	}

	return knapsackTDMemo(memo, profits, weights, capacity, 0)
}

func knapsackTDMemo(memo [][]int, profits, weights []int, capacity, currentIndex int) int {
	if capacity <= 0 || currentIndex >= len(profits) {
		return 0
	}

	// return immediately if found in cache.
	if memo[currentIndex][capacity] != 0 {
		return memo[currentIndex][capacity]
	}

	// calculate the profit for the item at the current index.
	profit1 := 0
	if weights[currentIndex] <= capacity {
		profit1 = profits[currentIndex] + knapsackTDMemo(memo, profits, weights, capacity-weights[currentIndex], currentIndex+1)
	}

	// process after excluding the item at the current index.
	profit2 := knapsackTDMemo(memo, profits, weights, capacity, currentIndex+1)

	return common.Max(profit1, profit2)
}

func knapsackBU(profits, weights []int, capacity int) int {
	n := len(profits)
	if capacity <= 0 || n == 0 || len(weights) != n {
		return 0
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, capacity+1)
	}

	// if we have only one weight, we will take it if it is not more than the
	// capacity.
	for i := 0; i < capacity+1; i++ {
		if weights[0] <= i {
			memo[0][i] = profits[0]
		}
	}

	for i := 1; i < n; i++ {
		for c := 1; c < capacity+1; c++ {
			profit1, profit2 := 0, 0

			// include the item if it's not bigger than the capacity.
			if weights[i] <= c {
				profit1 = profits[i] + memo[i-1][c-weights[i]]
			}

			// exclude the item.
			profit2 = memo[i-1][c]

			// take the maximum.
			memo[i][c] = common.Max(profit1, profit2)
		}
	}

	return memo[n-1][capacity]
}
