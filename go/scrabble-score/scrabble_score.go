package scrabble

import (
	"unicode"
)

//Score take a string in and return its scrabble value
func Score(s string) int {
	sum := 0
	for _, r := range s {
		switch unicode.ToLower(r) {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			sum++
		case 'd', 'g':
			sum += 2
		case 'b', 'c', 'm', 'p':
			sum += 3
		case 'f', 'h', 'v', 'w', 'y':
			sum += 4
		case 'k':
			sum += 5
		case 'j', 'x':
			sum += 8
		case 'q', 'z':
			sum += 10
		}
	}
	return sum
}
