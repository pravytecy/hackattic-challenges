package main

import (
	"fmt"
	"strings"
)

func compress(input string) {
	var sb strings.Builder
	count := 1
	for i := 1; i <= len(input); i++ {
		if i < len(input) && input[i] == input[i-1] {
			count++
		} else {
			if count > 2 {
				fmt.Fprintf(&sb, "%d%c", count, rune(input[i-1]))
			} else {
				sb.WriteString(strings.Repeat(string(input[i-1]), count))
			}
			count = 1
		}
	}
	fmt.Println(sb.String())
}
