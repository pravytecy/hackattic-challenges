package main

import (
	"fmt"
	"unicode"
)

func caseOfSnakes(input string) {
	var strippedInput string
	for i, char := range input {
		if char >= 65 && char <= 90 {
			strippedInput = input[i:]
			break
		}
	}
	output := ""
	for i, char := range strippedInput {
		if char >= 65 && char <= 90 && i == 0 {
			output += string(unicode.ToLower(rune(char)))

		} else if char >= 65 && char <= 90 {
			output += "_" + string(unicode.ToLower(rune(char)))
		} else {
			output += string(char)
		}
	}
	fmt.Println(output)
}
