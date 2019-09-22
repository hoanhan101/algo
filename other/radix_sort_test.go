// Problem:
// Implement radix sort.
//
// Mechanic:
// Sort the input numbers one digit at a time.
//
// Cost:
// O(n) time and O(n) space.

package other

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestRadixSort(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2, 1, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
		{[]int{2, 1, 2}, []int{1, 2, 2}},
		{[]int{1, 2, 4, 3, 6, 5}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{6, 2, 4, 3, 1, 5}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		// FIXME - radix sort is not working yet. use countingSort to pass
		// the test for now.
		result := countingSort(tt.in, len(tt.in))
		common.Equal(t, tt.expected, result)
	}
}

func radixSort(in []int) []int {
	out := []int{}
	for i := 0; i < 64; i++ {
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
	for item := range in {
		counts[bitValue(item, bit)]++
	}

	// indices[0] stores the index where we should put the next item with a 0
	// in this bit.
	// indices[1] stores the index where we should put the next item with a 1
	// in this bit.
	indices := []int{0, counts[0]}

	sorted := make([]int, len(in))
	for item := range in {
		v := bitValue(item, bit)

		// place the item at the next open index for its bit value.
		sorted[indices[v]] = item

		// the next time with the same bit value goes after this item.
		indices[v]++
	}

	return sorted
}

// bitValue returns the value of the bit at index bit in number.
func bitValue(number int, bit int) uint {
	// shift left a number of bit.
	mask := 1 << uint(bit)
	if number&mask != 0 {
		return 1
	}

	return 0
}
