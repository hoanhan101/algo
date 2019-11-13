/*
Problem:
- Determine whether an integer is a palindrome.

Assumption:
- Do this without extra space.
- Define negative integers as non-palindrome.

Example:
- Input: 101
  Output: true
- Input: 106
  Output: false

Approach:
- Use two-pointer approach where one starts at the first digit and one starts
  at the last digit, have them walk toward the middle and compare them at each
  step.

Solution:
- Calculate the division factor for the number to divide by to get its first digit.
- While the number is not equal to 0:
  - Get the first digit by diving the number by the division factor above.
  - Get the last digit by applying modulo the number by 10.
- Return false immediately if they are equal.
- Otherwise, chop of both digit by applying modulo the number by the division factor
  divide the result by 10.
- Update the division factor by dividing it by 100 since we already chopped of the first
  and last digits.

Cost:
- O(n) time, O(1) space.
*/

package leetcode

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestIsIntegerPalindrome(t *testing.T) {
	tests := []struct {
		in       int
		expected bool
	}{
		{-1, false},
		{0, true},
		{1, true},
		{10, false},
		{11, true},
		{101, true},
		{111, true},
		{110, false},
		{-101, false},
	}

	for _, tt := range tests {
		result := isIntegerPalindrome(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func isIntegerPalindrome(x int) bool {
	// return false immediately if the input is negative.
	if x < 0 {
		return false
	}

	// calculate the division factor for the number to divide by to get its
	// first digit.
	div := 1
	for x/div >= 10 {
		div *= 10
	}

	for x != 0 {
		l := x / div
		r := x % 10

		// return false immediately if two corresponding digits are the not
		// equal.
		if l != r {
			return false
		}

		// chop off the first and last digits.
		x = (x % div) / 10

		// update the division factor by dividing it by 100 since we already
		// chopped of the first and last digits.
		div /= 100
	}

	return true
}
