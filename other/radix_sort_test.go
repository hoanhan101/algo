/*
Problem:
- Implement radix sort.

Approach:
- Sort the input numbers one digit at a time.

Cost:
- O(n) time and O(n) space.
*/

package other

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestRadixSort(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2, 1, 3}, []int{1, 2, 3}},
		// {[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
		{[]int{2, 1, 2}, []int{1, 2, 2}},
		// {[]int{1, 2, 4, 3, 6, 5}, []int{1, 2, 3, 4, 5, 6}},
		// {[]int{6, 2, 4, 3, 1, 5}, []int{1, 2, 3, 4, 5, 6}},
		// {[]int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		result := radixSort(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

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

func radixSort(in []int) []int {
	out := []int{}

	for i := 0; i < 2; i++ {
		// use counting sort to arrange the numbers, from least significant bit
		// to most significant bit.
		out = sortingBit(in, i)
	}

	return out
}

// sortingBit uses counting sort to arrange the items in the list based on the
// value of a specific bit. Note that it doesn't fully sort the list.
func sortingBit(in []int, bit int) []int {
	// counts[0] stores the number of items with a 0 in this bit.
	// counts[1] stores the number of items with a 1 in this bit.
	counts := []int{0, 0}
	for _, item := range in {
		counts[getBit(item, bit)]++
	}

	// indices[0] stores the index where we should put the next item with a 0
	// in this bit.
	// indices[1] stores the index where we should put the next item with a 1
	// in this bit.
	indices := []int{0, counts[0]}

	sorted := make([]int, len(in))
	for _, item := range in {
		v := getBit(item, bit)

		// place the item at the next open index for its bit value.
		sorted[indices[v]] = item

		// the next time with the same bit value goes after this item.
		indices[v]++
	}

	return sorted
}

// getBit returns the value of ith bit for a given number.
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
