/*
read a person's graduation year from stdin and print out their age.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Print("Enter yo graduation year: ")

	// initialize a scanner to read from stdin.
	scanner := bufio.NewScanner(os.Stdin)

	// scan then print out the output.
	var in string
	if scanner.Scan() {
		in = scanner.Text()
	}

	err := scanner.Err()
	if err != nil {
		fmt.Println(err)
	}

	gradYear, err := strconv.Atoi(in)
	if err != nil {
		fmt.Println(err)
	}

	age := time.Now().Year() - (gradYear - 22)
	fmt.Println("Yo age is", age)
}
