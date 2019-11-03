/*
Problem:
- Given two numbers, m and n, and two bit positions, i and j, insert m into n
  such that m starts at bit j and ends at bit i.

Example:
- Input: n = 10000000000, m = 10011, i = 2, j = 6
  Output: n = 10001001100

Approach:
- Clear the bits j through i in n using a custom mask.
- Shift m so that it lines up with bits j through i.
- Merge them together.
*/

package ctci

import (
	"strconv"
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestInsertBit(t *testing.T) {
	tests := []struct {
		in1      int64
		in2      int64
		in3      int64
		in4      int64
		expected int64
	}{
		{
			binToInt("10000000000"),
			binToInt("10011"),
			0,
			4,
			binToInt("10000010011"),
		},
		{
			binToInt("10000000000"),
			binToInt("10011"),
			1,
			5,
			binToInt("10000100110"),
		},
		{
			binToInt("10000000000"),
			binToInt("10011"),
			2,
			6,
			binToInt("10001001100"),
		},
		{
			binToInt("10000000000"),
			binToInt("10011"),
			3,
			7,
			binToInt("10010011000"),
		},
		{
			binToInt("10000000000"),
			binToInt("10011"),
			4,
			8,
			binToInt("10100110000"),
		},
		{
			binToInt("10000000000"),
			binToInt("10011"),
			5,
			9,
			binToInt("11001100000"),
		},
		{
			binToInt("10000000000"),
			binToInt("10011"),
			6,
			10,
			binToInt("10011000000"),
		},
	}

	for _, tt := range tests {
		result := insertBit(tt.in1, tt.in2, tt.in3, tt.in4)
		common.Equal(t, tt.expected, result)
	}
}

func insertBit(n, m, i, j int64) int64 {
	// clear the bits i through j by creating a bitmask that have all 1s,
	// except for 0s in the bits i through j.
	ones := ^0                  // create a sequence of 1s.
	left := ones << uint(j+1)   // have 1s before position j, then 0s.
	right := 1<<uint(i) - 1     // have 1s after position i
	mask := int64(left | right) // merge left and right to create a mask of all 1s,except for 0s between i and j.

	// can clear the bits now after having the mask.
	nCleared := n & mask

	// shift m so that it lines up with bits j through i.
	mShifted := m << uint(i)

	// merge them together and return.
	return nCleared | mShifted
}

// binToInt parses a given binary representation into an integer.
func binToInt(s string) int64 {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return -1
	}

	return i
}
