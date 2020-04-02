/*
Problem:
- Given a set of positive numbers, find if we can partition it into two subsets
  such that the sum of elements in both the subsets is equal.

Example:
- Input: []int{1, 2, 3, 4}
  Output: true
  Explanation: The set can be partitioned into {1, 4} and {2, 3}
- Input: []int{1, 1, 3, 4, 7}
  Output: true
  Explanation: The set can be partitioned into {1, 3, 4} and {1, 7}
- Input: []int{2, 3, 4, 6}
  Output: false

Brute-force approach:
- Assume if s is the total sum of all numbers, the two equal subset must have a sum of s/2.
- Use Knapsack approach, create a set that includes i which does not exceed s/2 and a
  set that does not.

Cost:
- Brute-force: O(2^n) time, O(n) space.
- Top-down: O(n*s) time, O(n*s) space where n is the number of items, s is the total sum of numbers.
- Bottom-up: O(n*s) time, O(n*s) space where n is the number of items, s is the total sum of numbers.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestCanPartition(t *testing.T) {
	tests := []struct {
		in       []int
		expected bool
	}{
		{[]int{1, 2, 3, 4}, true},
		{[]int{1, 1, 3, 4, 7}, true},
		{[]int{2, 3, 4, 6}, false},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			canPartitionBF(tt.in),
		)

		common.Equal(
			t,
			tt.expected,
			canPartitionTD(tt.in),
		)

		common.Equal(
			t,
			tt.expected,
			canPartitionBU(tt.in),
		)
	}
}

func canPartitionBF(nums []int) bool {
	sum := common.SumInt(nums)

	// return false if the total sum is odd since we cannot have 2 subsets with equal sum.
	if sum%2 != 0 {
		return false
	}

	return canPartitionBFRecur(nums, sum/2, 0)
}

func canPartitionBFRecur(nums []int, sum, currentIndex int) bool {
	if sum == 0 {
		return true
	}

	n := len(nums)
	if n == 0 || currentIndex >= n {
		return false
	}

	if nums[currentIndex] <= sum {
		if canPartitionBFRecur(nums, sum-nums[currentIndex], currentIndex+1) {
			return true
		}
	}

	return canPartitionBFRecur(nums, sum, currentIndex+1)
}

func canPartitionTD(nums []int) bool {
	sum := common.SumInt(nums)

	if sum%2 != 0 {
		return false
	}

	// initialize a 2D array to act as a memoization table, in which
	// 1 is for true, 0 for false.
	memo := make([][]int, len(nums))
	for i := range memo {
		memo[i] = make([]int, sum/2+1)
	}

	return canPartitionTDMemo(memo, nums, sum/2, 0) == 1
}

func canPartitionTDMemo(memo [][]int, nums []int, sum, currentIndex int) int {
	if sum == 0 {
		return 1
	}

	n := len(nums)
	if n == 0 || currentIndex >= n {
		return 0
	}

	if nums[currentIndex] <= sum {
		if canPartitionTDMemo(memo, nums, sum-nums[currentIndex], currentIndex+1) == 1 {
			memo[currentIndex][sum] = 1
			return 1
		}
	}

	memo[currentIndex][sum] = canPartitionTDMemo(memo, nums, sum, currentIndex+1)

	return memo[currentIndex][sum]
}

// TODO - document the approach.
func canPartitionBU(nums []int) bool {
	sum := common.SumInt(nums)

	if sum%2 != 0 {
		return false
	}

	sum = sum / 2

	n := len(nums)
	tabu := make([][]bool, n)
	for i := range tabu {
		tabu[i] = make([]bool, sum+1)
	}

	for i := 0; i < n; i++ {
		tabu[i][0] = true
	}

	for i := 1; i < sum+1; i++ {
		tabu[0][i] = nums[0] == i
	}

	for i := 1; i < n; i++ {
		for j := 1; j < sum+1; j++ {
			if tabu[i-1][j] {
				tabu[i][j] = tabu[i-1][j]
			} else if j >= nums[i] {
				tabu[i][j] = tabu[i-1][j-nums[i]]
			}
		}
	}

	return tabu[n-1][sum]
}
