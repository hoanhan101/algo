/*
Problem:
- Write an algorithm to determine if a number is happy.
- Any number will be called a happy number if, after repeatedly replacing
  it with a number equal to the sum of the square of all of its digits,
  leads us to 1.

Example:
- Input: 19
  Output: true
  Explanation:
	  1^2 + 9^2 = 82
	  8^2 + 2^2 = 68
	  6^2 + 8^2 = 100
	  1^2 + 0^2 + 0^2 = 1

Approach:
- Since the process always end in a cycle, we can use a similar approach to
  finding a cycle in linked list problem.
- Once is cycle is found, check if it is stuck on 1.

Cost:
- O(n) time, O(1) space.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestIsHappy(t *testing.T) {
	tests := []struct {
		in       int
		expected bool
	}{
		{23, true},
		{19, true},
		{12, false},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			isHappy(tt.in),
		)
	}
}

func isHappy(num int) bool {
	slow, fast := num, num

	for {
		slow = squareSum(slow)
		fast = squareSum(squareSum(fast))

		if slow == fast {
			break
		}
	}

	return slow == 1
}

func squareSum(num int) int {
	sum := 0

	for num > 0 {
		lastDigit := num % 10
		sum += lastDigit * lastDigit

		// get rid of last digit.
		num /= 10
	}

	return sum
}
