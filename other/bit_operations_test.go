// Problem:
// Implement common bit operations such as Get Bit, Set Bit, Clear Bit and
// Update Bit.

package other

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestGetBit(t *testing.T) {
	tests := []struct {
		in1      int
		in2      int
		expected uint
	}{
		{6, 1, 1},
		{6, 2, 1},
		{6, 3, 0},
		{6, 4, 0},
		{6, 5, 0},
		{6, 6, 0},
		{6, 7, 0},
		{6, 8, 0},
	}

	for _, tt := range tests {
		result := getBit(tt.in1, tt.in2)
		common.Equal(t, tt.expected, result)
	}
}

// getBit returns the value of bit ith for a given number.
func getBit(number, i int) uint {
	// perform an AND with number to clear all bits other than the one at bit
	// i.
	mask := 1 << uint(i)

	// if the value is not 0, bit 1 must have a 1.
	if number&mask != 0 {
		return 1
	}

	return 0
}
