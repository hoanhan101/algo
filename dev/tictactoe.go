package main

import (
	"fmt"
)

var board [3][3]string

func main() {
	initBoard()
	place(0, 0, "X")
	place(1, 1, "X")
	place(2, 2, "X")
}

func initBoard() {
	for r := range board {
		for c := range board {
			board[r][c] = "_"
		}
	}

	printBoard()
}

func printBoard() {
	fmt.Println()
	for r := range board {
		for c := range board {
			fmt.Printf("%v ", board[r][c])
		}
		fmt.Println()
	}
}

func place(r, c int, s string) {
	board[r][c] = s

	printBoard()
}
