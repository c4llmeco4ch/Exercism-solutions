package grains

import (
	"errors"
)

const start uint64 = 1

//Total return the total number of grains for all spaces on the board
func Total(input int, wantSingle bool) (uint64, error) {
	if input <= 0 || input > 64 {
		return 0, errors.New("Invalid number entered")
	}
	top, pos := input, 1
	if wantSingle {
		pos = input
	}
	sum := uint64(0)
	for top >= pos {
		sum += (start << (pos - 1))
		pos++
	}
	return sum, nil
}
