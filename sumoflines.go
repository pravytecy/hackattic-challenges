package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumOfLines(input string) {
	val := strings.Fields(input)
	sum := 0

	for _, v := range val {
		switch {
		case strings.HasPrefix(v, "0x") || strings.HasPrefix(v, "0X"):
			value, _ := strconv.ParseInt(v, 0, 64)
			sum += int(value)

		case strings.HasPrefix(v, "0o") || strings.HasPrefix(v, "0O"):
			value, _ := strconv.ParseInt(v, 0, 64)
			sum += int(value)

		case strings.HasPrefix(v, "0b") || strings.HasPrefix(v, "0B"):
			value, _ := strconv.ParseInt(v, 0, 64)
			sum += int(value)

		default:
			value, err := strconv.ParseInt(v, 0, 64)
			if err == nil {
				sum += int(value)
			} else {
				sum += int(v[0])
			}

		}
	}
	fmt.Println(sum)
}
