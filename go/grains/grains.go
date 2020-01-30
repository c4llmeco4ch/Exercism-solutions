package grains

import (
	"errors"
)

const start uint64 = 1

//Square return the number of grains at space 'input'
func Square(input int) (uint64, error) {
	if input <= 0 || input > 64 {
		return 0, errors.New("Invalid number entered")
	}
	return start << (input - 1), nil
}

//Total return the total number of grains for all spaces on the board
func Total() uint64 {
	top, pos := 64, 1
	sum := uint64(0)
	for top >= pos {
		val, err := Square(pos)
		if err != nil {
			return 0 //Dealing with invald pos reached
		}
		sum += val
		pos++
	}
	return sum
}
