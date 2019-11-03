/*
Problem:
- Given an unsorted list scores (integer) and a highest possible score
  (integer), return a sorted list utilizing that fact.

Example:
- Input: []int{37, 89, 41, 65, 91, 53}, 100
  Output: []int{91, 89, 65, 53, 41, 37}

Approach:
- Utilize the highest score to allocate a fix-sized list ahead of time where
  where its indices represent the scores themselves and its values represent
  how many time these scores appear in the list.

Solution:
- Utilize the highest score to allocate a fix-sized list ahead of time.
- Iterate through the unsorted list and populate the scores where indices
  represent scores and values represent how many time the score appears.
- Iterate backward and add the score to the output list so that it is
  sorted from highest to lowest.

Cost:
- O(n) time, O(n) space.
*/

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestSortScores(t *testing.T) {
	tests := []struct {
		in1      []int
		in2      int
		expected []int
	}{
		{[]int{}, 0, []int{}},
		{[]int{37, 89, 41, 65, 91, 53}, 100, []int{91, 89, 65, 53, 41, 37}},
		{[]int{37, 89, 41, 65, 91, 53, 65}, 100, []int{91, 89, 65, 65, 53, 41, 37}},
	}
	for _, tt := range tests {
		result := sortScores(tt.in1, tt.in2)
		common.Equal(t, tt.expected, result)
	}
}

func sortScores(scores []int, highestPossibleScore int) []int {
	// utilize highest score to create a fix-sized list of score counts.
	counts := make([]int, highestPossibleScore+1)
	out := make([]int, 0)

	// populate the scores where indices represent scores and values represent
	// how many time the score appears.
	for _, score := range scores {
		counts[score]++
	}

	// iterate backward and add the score to the output list so that it is
	// sorted from highest to lowest.
	for i := len(counts) - 1; i > -1; i-- {
		count := counts[i]

		for j := 0; j < count; j++ {
			out = append(out, i)
		}
	}

	return out
}
