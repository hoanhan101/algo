package common

import (
	"math"
	"math/rand"
)

// Mimax returns min and max from a list of numbers.
func Mimax(nums ...int) (int, int) {
	min, max := nums[0], nums[0]

	for _, num := range nums {
		if min > num {
			min = num
		}

		if max < num {
			max = num
		}
	}

	return min, max
}

// Random rertuns a random number over a range
func Random(min, max int) int {
	if min == max {
		return min
	}

	return rand.Intn(max-min) + min
}

// ChanToSlice pushes values from a channel to a slice.
func ChanToSlice(ch chan int) []int {
	out := []int{}

	for v := range ch {
		out = append(out, v)
	}

	return out
}

// Contain checks if the target is in a slice.
func Contain(s []int, target int) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}

	return false
}

// IsMoreThan1Apart checks if two integers are more than 1 apart.
func IsMoreThan1Apart(a, b int) bool {
	if math.Abs(float64(a)-float64(b)) > 1 {
		return true
	}

	return false
}
