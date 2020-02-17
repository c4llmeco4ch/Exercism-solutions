package grains

import (
	"errors"
)

const (
	start uint64 = 1
	max   uint64 = 64
)

//Square return the number of grains at space 'input'
func Square(input int) (uint64, error) {
	if input <= 0 || input > 64 {
		return 0, errors.New("Invalid number entered")
	}
	return 1 << (input - 1), nil
}

//Total return the total number of grains for all spaces on the board
func Total() uint64 {
	return (1 << 64) - 1
}
