/*
Problem:
- Given a matrix of m rows x n columns, return all elements of the matrix in spiral
  order.

Example:
- Input: [][]int{
	[]int{1, 2, 3},
	[]int{4, 5, 6},
	[]int{7, 8, 9},
  }
  Output: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}

Approach:
- Keep track of the current position as we traverse the matrix in spiral order.

Cost:
- O(n) time, O(n) space.
*/

package leetcode

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		in       [][]int
		expected []int
	}{
		{
			[][]int{},
			[]int{},
		},
		{
			[][]int{
				[]int{1},
			},
			[]int{1},
		},
		{
			[][]int{
				[]int{1, 2},
			},
			[]int{1, 2},
		},
		{
			[][]int{
				[]int{1},
				[]int{2},
			},
			[]int{1, 2},
		},
		{
			[][]int{
				[]int{1, 2},
				[]int{3, 4},
			},
			[]int{1, 2, 4, 3},
		},
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
			},
			[]int{1, 2, 3, 6, 5, 4},
		},
		{
			[][]int{
				[]int{1, 2},
				[]int{3, 4},
				[]int{5, 6},
			},
			[]int{1, 2, 4, 6, 5, 3},
		},
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{
				[]int{1, 2, 3, 4},
				[]int{5, 6, 7, 8},
				[]int{9, 10, 11, 12},
			},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
				[]int{10, 11, 12},
			},
			[]int{1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8},
		},
		{
			[][]int{
				[]int{1, 2, 3, 4},
				[]int{5, 6, 7, 8},
				[]int{9, 10, 11, 12},
				[]int{13, 14, 15, 16},
			},
			[]int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
		},
	}

	for _, tt := range tests {
		common.Equal(
			t,
			tt.expected,
			spiralOrder(tt.in),
		)
	}
}

func spiralOrder(matrix [][]int) []int {
	out := []int{}
	if len(matrix) == 0 {
		return out
	}

	// initialize m and n to be the number of rows and columns.
	m := len(matrix)
	n := len(matrix[0])

	// row and col keeps track of the current row and column position as we
	// traverse the matrix.
	row := 0
	col := -1

	for {
		// add all values in the row by keeping the row position and increasing
		// the col position.
		for i := 0; i < n; i++ {
			col++
			out = append(out, matrix[row][col])
		}

		// decrease the number of rows so that we don't have to revisit it again.
		// if it is 0, it means we have added all the values, return immediately.
		m--
		if m == 0 {
			break
		}

		// add all values in the column by keeping the column position and increasing
		// the row position.
		for i := 0; i < m; i++ {
			row++
			out = append(out, matrix[row][col])
		}

		// decrease the number of columns so that we don't have to revisit it again.
		n--
		if n == 0 {
			break
		}

		// add all values in the row by keeping the row position and decreasing
		// the col position.
		for i := 0; i < n; i++ {
			col--
			out = append(out, matrix[row][col])
		}

		// decrease the number of rows so that we don't have to revisit it again.
		m--
		if m == 0 {
			break
		}

		// add all values in the column by keeping the column position and decreasing
		// the row position.
		for i := 0; i < m; i++ {
			row--
			out = append(out, matrix[row][col])
		}

		// decrease the number of columns so that we don't have to revisit it again.
		n--
		if n == 0 {
			break
		}
	}

	return out
}
