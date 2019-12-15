/*
	count.go counts the number of files in given directories.
	Since one file contains one problem, the program's output can also
	tell how many questions that we have completed so far.
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/tabwriter"
)

func main() {
	// format in tab-separated columns with a tab stop of 8.
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	total := 0
	directories := []string{"ctci", "gtci", "interviewcake", "leetcode", "other"}

	for _, dir := range directories {
		n := countFiles(dir)
		fmt.Fprintf(w, "%v\t%v\n", dir, n)

		total += n
	}

	fmt.Fprintf(w, "-------------\t--\n")
	fmt.Fprintf(w, "total\t%v\n", total)
	w.Flush()
}

func countFiles(dir string) int {
	counter := 0

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if dir != path {
				counter++
			}
			return nil
		})

	if err != nil {
		fmt.Println(err)
	}

	return counter
}
