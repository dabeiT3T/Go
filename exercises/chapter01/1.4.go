package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	files := os.Args[1:]

	for _, filename := range files {

		counts := make(map[string]int)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}

		input := bufio.NewScanner(f)
		for input.Scan() {
			text := input.Text()
			counts[text]++
			// if duplicated, print the filename
			if counts[text] > 1 {
				fmt.Println(filename)
				break
			}
		}

		f.Close()
	}
}
