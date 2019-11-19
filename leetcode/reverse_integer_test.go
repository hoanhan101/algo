/*
Problem:
- Given a 64-bit integer, reverse its digits.

Assumption:
- Negative numbers are also valid.
- Must handle the case where the reversed integer is overflow.

Example:
- Input: 123
  Output: 321
- Input: -123
  Output: -321
- Input: 8085774586302733229 Output: 0
  Explanation: The reversed integer 9223372036854775808 overflows by 1 so we return 0.

Approach:
- Use modulo by 10 to get a digit at ones' place of the input and
  dividing by 10 to shift it to the right (eliminate the ones' place).

Solution:
- Initialize the output as a 64-bit integer.
- If the input is not zero yet, keep multiply the output and diving the
  input by 10 so that: output = (output * 10) + (input % 10) and
  input = input / 10.
- Also, remember to check the overflow/underflow of the output and return
  0 if necessary.

Cost:
- O(m) time, O(1) space, where m is log10 of the input.
*/

package leetcode

import (
	"math"
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestReverseInteger(t *testing.T) {
	tests := []struct {
		in       int64
		expected int64
	}{
		{0, 0},
		{1, 1},
		{10, 1},
		{12, 21},
		{123, 321},
		{-123, -321},
		{8085774586302733229, 0},
	}

	for _, tt := range tests {
		result := reverseInteger(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func reverseInteger(in int64) int64 {
	var out int64

	for in != 0 {
		// check for overflow/underflow before multiplying by 10.
		// if common.Abs(int(out)) > math.MaxInt64/10 {
		if common.Abs(int(out)) > math.MaxInt64/10 {
			return 0
		}

		out = out*10 + in%10
		in /= 10
	}

	return out
}
