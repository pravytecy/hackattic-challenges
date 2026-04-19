package main

import "fmt"

func day(input int) {
	var DAYS = [7]string{"Thursday", "Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday"}
	remainder := int(input) % 7
	if remainder < 0 {
		remainder += 7
	}
	fmt.Println(DAYS[remainder])
}
