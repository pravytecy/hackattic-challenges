package main

import "fmt"

func paranthesis(input string) {
	count := 0
	for _, ch := range input {
		if ch == '(' {
			count++
		} else {
			count--
		}
		if count < 0 {
			fmt.Println("false")
			return
		}
	}
	if count == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
