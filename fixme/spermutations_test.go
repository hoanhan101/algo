/*
Problem:
- Given a set of distinct

Example:
- Input: [1 2]
  Output: [[] [1] [2] [1 2]]
- Input: [1 2 5]
  Output: [[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]

Approach:
- Similar to the previous problem, but we do two more steps.
- First, sort the set so that duplicates are next to each other.
- Second, when we encounter a duplicate while iterating the set,
  only create subsets from the subsets that added previously.
- Can use a two-pointer approach to update their start and end window
  accordingly.

Cost:
- O(2^n) time, O(2^n) space since we would have a total of 2^n subsets.
*/

package fixme

/*
import (
	"container/list"
	"sort"
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFindSPermutations(t *testing.T) {
	tests := []struct {
		in       []int
		expected [][]int
	}{
		{[]int{}, [][]int{{}}},
		{[]int{1}, [][]int{{}, {1}}},
		{[]int{1, 2}, [][]int{{}, {1}, {2}, {1, 2}}},
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			findSPermutations(tt.in),
		)
	}
}

func findSPermutations(nums []int) [][]int {
	subsets := [][]int{}
	permutations := list.New()
	permutations.PushBack([]int{})

	for _, num := range nums {
		// take all existing permutations to add the current number to.
		for e := permutations.Front(); e != nil; e = e.Next() {
			// pop the front of the list.
			existing := permutations.Front()
			permutations.Remove(existing)

			// add the current number at every position of the existing one.
			for i := 0; i < len(existing.Value.([]int)); i++ {
				newp := existing
				// FIXME - InsertAfter() inserts a new element after an
				// element, not after an index.
				// newp.InsertAfter()
				if len(newp.Value.([]int)) == len(nums) {
					subsets = append(subsets, newp.Value.([]int))
				} else {
					permutations.PushBack(newp)
				}
			}
		}
	}

	return subsets
}
*/
