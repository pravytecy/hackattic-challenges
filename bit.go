package main

func encodeBits(input string) uint16 {
	var result uint16 = 0
	for _, ch := range input {
		result = result << 1
		switch ch {
		case '#':
			result = result | 1
		}
	}
	return result
}
