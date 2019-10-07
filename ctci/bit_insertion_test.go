// Problem:
// Given two numbers, m and n, and two bit positions, i and j, insert m into n
// such that m starts at bit j and ends at bit i.
//
// Example:
// Input: n = 10000000000, m = 10011, i = 2, j = 6
// Output: n = 10001001100
//
// Solution:
// Clear the bits j through i in n using a custom mask.
// Shift m so that it lines up with bits j through i.
// Merge them together.

package ctci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestInsertBit(t *testing.T) {
	tests := []struct {
		in1      int
		in2      int
		in3      int
		in4      int
		expected int
	}{
		{
			1024, // n = 10000000000
			19,   // m = 10011
			2,    // i = 2
			6,    // j = 6
			1100, // n = 10001001100
		},
	}

	for _, tt := range tests {
		result := insertBit(tt.in1, tt.in2, tt.in3, tt.in4)
		common.Equal(t, tt.expected, result)
	}
}

func insertBit(n, m, i, j int) int {
	// clear the bits i through j by creating a bitmask that have all 1s,
	// except for 0s in the bits i through j.
	ones := ^0                // create a sequence of 1s.
	left := ones << uint(j+1) // have 1s before position j, then 0s.
	right := 1<<uint(i) - 1   // have 1s after position i
	mask := left | right      // merge left and right to create a mask of all 1s,except for 0s between i and j.

	// can clear the bits now after having the mask.
	nCleared := n & mask

	// shift m so that it lines up with bits j through i.
	mShifted := m << uint(i)

	// merge them together and return.
	return nCleared | mShifted
}
