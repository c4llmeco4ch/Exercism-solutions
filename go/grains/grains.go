package grains

import (
	"errors"
)

var (
	board = [64]uint64{0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	}
)

//Square return the number of grains at space 'input'
func Square(input int) (uint64, error) {
	if input <= 0 || input > 64 {
		return 0, errors.New("Invalid input number")
	}
	if input == 1 {
		board[0] = 1
		return 1, nil
	} else if board[input-1] != 0 {
		return board[input-1], nil
	}
	s, _ := Square(input - 1)
	board[input-1] = s * 2
	return s * 2, nil
}

//Total return the total number of grains for all spaces on the board
func Total() uint64 {
	top, pos := 64, 1
	sum := uint64(0)
	for top >= pos {
		val, _ := Square(pos)
		sum += val
		pos++
	}
	return sum
}
