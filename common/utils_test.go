package common

import (
	"reflect"
	"testing"
)

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
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
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
		result := Random(tt.in1, tt.in2)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
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
		result := Contain(tt.s, tt.target)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
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
		result := IsMoreThan1Apart(tt.in1, tt.in2)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
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
		result := IsLessThan1Apart(tt.in1, tt.in2)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}
