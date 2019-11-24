/*
read from stdin.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// initialize a scanner to read from stdin.
	scanner := bufio.NewScanner(os.Stdin)

	// scan then print out the output.
	if scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err := scanner.Err()
	if err != nil {
		fmt.Println(err)
	}
}
