/*
Problem:
- Given an array containing n objects where each object, when created,
  was assigned a unique number from 1 to n based on their creation sequence.
  This means that the object with sequence number 3 was created just before
  the object with sequence number 4.
- Write a function to sort the objects in-place on their creation sequence
  number in O(n) and without any extra space.

Example:
- Input: []int{6, 3, 5, 2, 4, 1}
  Output: []int{1, 2, 3, 4, 5, 6}

Approach:
- Use the fact that we are given a range of 1-n, can try placing each number at
  its correct index in the array, say 1 at 0, 2 at 1, 3 at 2 and so on.
- Iterate through the array and if the current number is not at the correct index,
  swap it with the number at its correct index.

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestCyclicSort(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{6, 4, 5, 2, 3, 1}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		cyclicSort(tt.in)
		common.Equal(
			t,
			tt.expected,
			tt.in,
		)
	}
}

func cyclicSort(nums []int) {
	i := 0
	for i < len(nums) {
		// if the current number is not at the correct index, swap it with the
		// one that is at the correct index.
		correctIndex := nums[i] - 1
		if nums[i] != nums[correctIndex] {
			common.Swap(nums, i, correctIndex)
		} else {
			i++
		}
	}
}
