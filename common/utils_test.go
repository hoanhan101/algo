package common

import (
	"testing"
)

func TestSwapInt(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 1, 3, 4, 5}},
		{[]int{2, 1, 3, 4, 5}, []int{2, 3, 1, 4, 5}},
		{[]int{2, 3, 1, 4, 5}, []int{2, 3, 4, 1, 5}},
		{[]int{2, 3, 4, 1, 5}, []int{2, 3, 4, 5, 1}},
	}

	for i, tt := range tests {
		SwapInt(tt.in, i, i+1)
		Equal(t, tt.expected, tt.in)
	}
}

func TestSwapString(t *testing.T) {
	tests := []struct {
		in       []string
		expected []string
	}{
		{[]string{"a", "b", "c", "d", "e"}, []string{"b", "a", "c", "d", "e"}},
		{[]string{"b", "a", "c", "d", "e"}, []string{"b", "c", "a", "d", "e"}},
		{[]string{"b", "c", "a", "d", "e"}, []string{"b", "c", "d", "a", "e"}},
		{[]string{"b", "c", "d", "a", "e"}, []string{"b", "c", "d", "e", "a"}},
	}

	for i, tt := range tests {
		SwapString(tt.in, i, i+1)
		Equal(t, tt.expected, tt.in)
	}
}

func TestMimax(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 5}},
		{[]int{-1, -2, -3, -4, -5}, []int{-5, -1}},
		{[]int{-2, -1, 0, 1, 2}, []int{-2, 2}},
		{[]int{-10, -10, 1, 3, 2}, []int{-10, 3}},
		{[]int{1, 10, -5, 1, -100}, []int{-100, 10}},
	}

	for _, tt := range tests {
		min, max := Mimax(tt.in...)
		result := []int{min, max}
		Equal(t, tt.expected, result)
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{-1, -2, -3, -4, -5}, -5},
		{[]int{-2, -1, 0, 1, 2}, -2},
		{[]int{-10, -10, 1, 3, 2}, -10},
		{[]int{1, 10, -5, 1, -100}, -100},
	}

	for _, tt := range tests {
		Equal(t, tt.expected, Min(tt.in...))
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{-1, -2, -3, -4, -5}, -1},
		{[]int{-2, -1, 0, 1, 2}, 2},
		{[]int{-10, -10, 1, 3, 2}, 3},
		{[]int{1, 10, -5, 1, -100}, 10},
	}

	for _, tt := range tests {
		Equal(t, tt.expected, Max(tt.in...))
	}
}

func TestRandom(t *testing.T) {
	tests := []struct {
		in1      int
		in2      int
		expected int
	}{
		{1, 10, 6},
		{-10, 10, -3},
		{0, 10, 7},
		{4, 4, 4},
	}

	for _, tt := range tests {
		Equal(t, tt.expected, Random(tt.in1, tt.in2))
	}
}

func TestContain(t *testing.T) {
	tests := []struct {
		s        []int
		target   int
		expected bool
	}{
		{[]int{}, 0, false},
		{[]int{1}, 0, false},
		{[]int{1}, 1, true},
		{[]int{0, 1, 2, 3, 101}, 2, true},
	}

	for _, tt := range tests {
		Equal(t, tt.expected, Contain(tt.s, tt.target))
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		in       int
		expected int
	}{
		{0, 0},
		{1, 1},
		{-1, 1},
	}

	for _, tt := range tests {
		Equal(t, tt.expected, Abs(tt.in))
	}
}

func TestAbsDiff(t *testing.T) {
	tests := []struct {
		in1      int
		in2      int
		expected int
	}{
		{1, 1, 0},
		{1, 2, 1},
		{1, 3, 2},
	}

	for _, tt := range tests {
		Equal(t, tt.expected, AbsDiff(tt.in1, tt.in2))
	}
}

func TestIsMoreThan1Apart(t *testing.T) {
	tests := []struct {
		in1      int
		in2      int
		expected bool
	}{
		{1, 1, false},
		{1, 2, false},
		{1, 3, true},
	}

	for _, tt := range tests {
		Equal(t, tt.expected, IsMoreThan1Apart(tt.in1, tt.in2))
	}
}

func TestIsLessThan1Apart(t *testing.T) {
	tests := []struct {
		in1      int
		in2      int
		expected bool
	}{
		{1, 1, true},
		{1, 2, true},
		{1, 3, false},
	}

	for _, tt := range tests {
		Equal(t, tt.expected, IsLessThan1Apart(tt.in1, tt.in2))
	}
}
