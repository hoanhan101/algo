/*
Problem:
- Given a list of integers, return a corresponding list where every index
  holds the product of every other values except the value in that index.
  And, you can't use division!

Example:
- Input: []int{1, 7, 3, 4}
  Output: []int{84, 12, 28, 21}

Approach:
- Iterate through the list and at each step, calculate the product of all
  the integers before each index and the product of all the integers after
  each index.

Solution:
- Iterate through the list from left to right and get the product of all
  other numbers before their indices.
- Iterate through the list from right to left and get the product of all
  other numbers after their indices.

Example walkthrough:
- After the first iteration, we end up with:
  []int{1*1, 1*1, 1*7, 1*7*3} or []int{1, 1, 7, 21}
- After the second iteration, we end up with:
  []int{1*4*3*7, 1*4*3, 7*4, 21*1} or []{84, 12, 28, 21}

Cost:
- O(n) time, O(n) space.
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestGetProductOfOthers(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 7, 3, 4}, []int{84, 12, 28, 21}},
		{[]int{2, 4, 10}, []int{40, 20, 8}},
		{[]int{2, 4, 0}, []int{0, 0, 8}},
	}

	for _, tt := range tests {
		result := getProductOfOthers(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func getProductOfOthers(list []int) []int {
	if len(list) < 2 {
		return []int{}
	}

	out := make([]int, len(list))

	// get product of all other numbers before their indices.
	start1 := 1
	for i := 0; i < len(list); i++ {
		out[i] = start1
		start1 *= list[i]
	}

	// get product of all other numbers after their indices then multiply them
	// with their corresponding values.
	start2 := 1
	for i := len(list) - 1; i > -1; i-- {
		out[i] *= start2
		start2 *= list[i]
	}

	return out
}
