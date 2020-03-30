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
- Input:
	weight :   1   |    6   |    10   |   16
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
- Top-down: TODO
- Bottom-up: TODO
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
		{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 7, 22},
		{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 6, 17},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			knapsackBF(tt.in1, tt.in2, tt.in3),
		)

		// common.Equal(
		// 	t,
		// 	tt.expected,
		// 	knapsackTD(tt.in),
		// )

		// common.Equal(
		// 	t,
		// 	tt.expected,
		// 	knapsackBU(tt.in),
		// )
	}
}

func knapsackBF(profits, weights []int, capacity int) int {
	return knapsackBFRecursive(profits, weights, capacity, 0)
}

func knapsackBFRecursive(profits, weights []int, capacity, currentIndex int) int {
	if capacity <= 0 || currentIndex >= len(profits) {
		return 0
	}

	// first, calculate the profit for the item at the current index.
	// if the total weight does not exceed the capacity, recursively process
	// the remaining capacity and items.
	profit1 := 0
	if weights[currentIndex] <= capacity {
		profit1 = profits[currentIndex] + knapsackBFRecursive(profits, weights, capacity-weights[currentIndex], currentIndex+1)
	}

	// second, recursively process after excluding the item at the current index.
	profit2 := knapsackBFRecursive(profits, weights, capacity, currentIndex+1)

	return common.Max(profit1, profit2)
}
