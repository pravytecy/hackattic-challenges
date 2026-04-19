package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid number:", err)
	}
	day(num)
	//fmt.Println(result)
}
