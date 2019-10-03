// Problem:
// Implement common Bit operations such as:
// - Get Bit
// - Set Bit
// - Clear Bit (from most significant bit to i and from i to the least significant bit)
// - Update Bit

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
		{6, 0, 0},
		{6, 1, 1},
		{6, 2, 1},
		{6, 3, 0},
		{6, 4, 0},
		{6, 5, 0},
		{6, 6, 0},
	}

	for _, tt := range tests {
		result := getBit(tt.in1, tt.in2)
		common.Equal(t, tt.expected, result)
	}
}

func TestSetBit(t *testing.T) {
	tests := []struct {
		in1      int
		in2      int
		expected int
	}{
		{6, 0, 7},
		{6, 1, 6},
		{6, 2, 6},
		{6, 3, 14},
		{6, 4, 22},
		{6, 5, 38},
		{6, 6, 70},
	}

	for _, tt := range tests {
		result := setBit(tt.in1, tt.in2)
		common.Equal(t, tt.expected, result)
	}
}

func TestClearBit(t *testing.T) {
	tests := []struct {
		in1      int
		in2      int
		expected int
	}{
		{6, 0, 6},
		{6, 1, 4},
		{6, 2, 2},
		{6, 3, 6},
		{6, 4, 6},
		{6, 5, 6},
		{6, 6, 6},
	}

	for _, tt := range tests {
		result := clearBit(tt.in1, tt.in2)
		common.Equal(t, tt.expected, result)
	}
}

func TestClearBitMSBToI(t *testing.T) {
	tests := []struct {
		in1      int
		in2      int
		expected int
	}{
		{6, 0, 0},
		{6, 1, 0},
		{6, 2, 2},
		{6, 3, 6},
		{6, 4, 6},
		{6, 5, 6},
		{6, 6, 6},
	}

	for _, tt := range tests {
		result := clearBitMSBToI(tt.in1, tt.in2)
		common.Equal(t, tt.expected, result)
	}
}

func TestClearBitIToLSB(t *testing.T) {
	tests := []struct {
		in1      int
		in2      int
		expected int
	}{
		{6, 0, 6},
		{6, 1, 4},
		{6, 2, 0},
		{6, 3, 0},
		{6, 4, 0},
		{6, 5, 0},
		{6, 6, 0},
	}

	for _, tt := range tests {
		result := clearBitIToLSB(tt.in1, tt.in2)
		common.Equal(t, tt.expected, result)
	}
}

// getBit returns the value of bit ith for a given number.
func getBit(number, i int) uint {
	// shift 1 over by i bits, creating a bitmask value.
	mask := 1 << uint(i)

	// perform an AND with number to clear all bits other than the one at bit i.
	// if the value is not 0, bit 1 must have a 1.
	if number&mask != 0 {
		return 1
	}

	return 0
}

// setBit sets the value of bit ith for a given number.
func setBit(number, i int) int {
	// shift 1 over by i bits, creating a bitmask value.
	mask := 1 << uint(i)

	// perform an OR with number to change the number at bit i without
	// affecting other values.,
	return number | mask
}

// clearBit clears the value of bit ith for a given number.
func clearBit(number, i int) int {
	// shift 1 over by i bits, creating a bitmask value and negate it.
	mask := ^(1 << uint(i))

	// perform an AND with number to clear out all the bits 1.
	return number & mask
}

// clearBitMSBToI clears all bits from the most significant bit to i for a
// given number.
func clearBitMSBToI(number, i int) int {
	// shift 1 over by i bits, creating a bitmask value and subtract 1 from it
	// to get a sequence of 0s followed by i 1s.
	mask := 1<<uint(i) - 1

	// perform an AND with number to leave the last i bits.
	return number & mask
}

// clearBitIToLSB clears all bits from bit i to the least significant bit
// for a given number.
func clearBitIToLSB(number, i int) int {
	// shift 1 over by i + 1 bits from a sequence of all 1s(which is -1) to get
	// a sequence of 1s followed by i 0 bits.
	mask := -1 << uint(i+1)

	// perform an AND with number to cancel out the rest of 1 bits.
	return number & mask
}
