package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	var scanner *bufio.Scanner

	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		scanner = bufio.NewScanner(f)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	for scanner.Scan() {
		input := scanner.Text()
		jsonSorting(input)
	}
	PrintSorted()
}
