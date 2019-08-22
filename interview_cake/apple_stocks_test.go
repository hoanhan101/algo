// Problem:
// Given a list of stock prices (integer) in chronological order, return the
// max profit from buying at earlier time and selling at later time.
//
// Example:
// Given:  []int{10, 7, 5, 8, 11, 9}
// Return: 6
//
// Solution:
// Use a greedy approach to keep a minimum and maximum price while iterating
// through the list.
//
// Cost:
// O(1) time, O(n) space.

package main

import (
	"reflect"
	"testing"
)

func TestGetMaxProfit(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{10}, 0},
		{[]int{10, 10}, 0},
		{[]int{10, 7, 5, 8, 11, 9}, 6},
		{[]int{10, 2, 5, 4, 7, 1}, 5},
		{[]int{10, 7, 2, 1}, -1},
	}

	for _, tt := range tests {
		result := getMaxProfit(tt.in)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

func getMaxProfit(stocks []int) int {
	// return 0 if the input is invalid
	if len(stocks) <= 2 {
		return 0
	}

	// initialize minPrice to be the first price in the list and maxProfit to
	// be the first possible profit that we could trade.
	minPrice := stocks[0]
	maxProfit := stocks[1] - stocks[0]

	for i := 1; i < len(stocks); i++ {
		currentPrice := stocks[i]
		potentialProfit := currentPrice - minPrice
		_, maxProfit = mimax(potentialProfit, maxProfit)
		minPrice, _ = mimax(minPrice, currentPrice)
	}

	return maxProfit
}
