package main

import (
	"math"
	"reflect"
	"testing"
)

func TestGetMaxProfit(t *testing.T) {
	tests := []struct {
		in       []float64
		expected float64
	}{
		{[]float64{}, 0},
		{[]float64{10}, 0},
		{[]float64{10, 10}, 0},
		{[]float64{10, 7, 5, 8, 11, 9}, 6},
		{[]float64{10, 2, 5, 4, 7, 1}, 5},
		{[]float64{10, 7, 2, 1}, -1},
	}

	for _, tt := range tests {
		result := getMaxProfit(tt.in)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

// getMaxProfit returns the max profit from a list of stock price from buying
// at earlier time and selling at later time.
func getMaxProfit(stocks []float64) float64 {
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
		maxProfit = math.Max(potentialProfit, maxProfit)
		minPrice = math.Min(minPrice, currentPrice)
	}

	return maxProfit
}
