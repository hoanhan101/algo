package common

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"testing"
)

// Swap two array values given their indices.
func Swap(v interface{}, i, j int) {
	switch a := v.(type) {
	case []int:
		SwapInt(a, i, j)
	case []string:
		SwapString(a, i, j)
	}
}

// SwapInt two array values given their indices.
func SwapInt(a []int, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

// SwapString two array values given their indices.
func SwapString(a []string, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

// Mimax returns min and max from a list of integers.
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

// Min returns min from a list of integers.
func Min(nums ...int) int {
	min := nums[0]

	for _, num := range nums {
		if min > num {
			min = num
		}
	}

	return min
}

// Max returns max from a list of integers.
func Max(nums ...int) int {
	max := nums[0]

	for _, num := range nums {
		if max < num {
			max = num
		}
	}

	return max
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

// ContainString checks if the target is in a slice.
func ContainString(s []string, target string) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}

	return false
}

// Abs returns the absolute value for a given integer.
func Abs(a int) int {
	return int(math.Abs(float64(a)))
}

// AbsDiff returns the absolute value of the difference between two integers.
func AbsDiff(a, b int) int {
	return Abs(a - b)
}

// IsMoreThan1Apart checks if two integers are more than 1 apart.
func IsMoreThan1Apart(a, b int) bool {
	return AbsDiff(a, b) > 1
}

// IsLessThan1Apart checks if two integers are less or equal than 1 apart.
func IsLessThan1Apart(a, b int) bool {
	return AbsDiff(a, b) <= 1
}

// Log prints out the map of logging context and value.
func Log(m map[string]interface{}) {
	fmt.Println("[debug] →")
	for k, v := range m {
		fmt.Printf("\t%v: %+v\n", k, v)
	}
	fmt.Println("[debug] □")
}

// Equal checks if two input are deeply equal.
func Equal(t *testing.T, expected, result interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("should be %v instead of %v", expected, result)
	}
}
