package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) {
	inputs := bufio.NewScanner(f)

	for inputs.Scan() {
		counts[inputs.Text()]++
	}
}

func readFromFile(counts map[string]int) error {
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			defer f.Close()
			if err != nil {
				return err
			}
			countLines(f, counts)

		}
	}
	return nil
}

func main() {
	counts := make(map[string]int)

	readFromFile(counts)

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
