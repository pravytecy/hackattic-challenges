package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//input := "#.#.#.###.#.##.#"
	var result uint16 = 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	for _, ch := range input {
		result = result << 1
		switch ch {
		case '#':
			result = result | 1
		}
	}
	fmt.Println(result)
}
