/*
Problem:
- Given an array of characters where each character represents a fruit tree,
  you are given two baskets and your goal is to put maximum number of fruits
  in each basket.

Constraints:
- Each basket can have only one type of fruit.
- You can start with any tree, but once you have started you can’t skip a tree.
  You will pick one fruit from each tree until you cannot, i.e., you will stop
  when you have to pick from a third fruit type.

Example:
- Input: fruits=["apple", "orange", "coconut", "apple", "coconut"]
  Output: 3
  Explanation: Can put 2 "cocunut" in 1 basket and 1 "apple" in other from
  subarray ["coconut", "apple", "coconut"]

Approach:
- Similar to "longest substring with k distinct characters" with k=2.

Cost:
- O(n) time, O(k) space where k is the number of characters in the map.
*/

package gtci

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFruitsIntoBaskets(t *testing.T) {
	tests := []struct {
		in       []string
		expected int
	}{
		{[]string{"apple", "orange", "coconut", "apple", "coconut"}, 3},
		{[]string{"apple", "orange", "coconut", "orange", "coconut", "orange", "coconut"}, 6},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			fruitsIntoBaskets(tt.in),
		)
	}
}

func fruitsIntoBaskets(fruits []string) int {
	maxLength, start := 0, 0

	// fruitsMap keeps track of fruits' frequencies.
	fruitsMap := map[string]int{}

	for end := range fruits {
		// insert fruits until we have 2 distinct fruits.
		endFruit := fruits[end]
		if _, ok := fruitsMap[endFruit]; !ok {
			fruitsMap[endFruit] = 0
		}
		fruitsMap[endFruit]++

		// shrink the window until there is no more than 2 distinct fruits.
		for len(fruitsMap) > 2 {
			startFruit := fruits[start]

			// decrement the frequency of the one going out of the window and
			// remove if its frequency is zero.
			fruitsMap[startFruit]--
			if fruitsMap[startFruit] == 0 {
				delete(fruitsMap, startFruit)
			}

			// increase the start index to move the window ahead by one element.
			start++
		}

		// update the maximum length at each step.
		maxLength = common.Max(maxLength, end-start+1)
	}

	return maxLength
}
