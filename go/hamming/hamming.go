//Package hamming Solutions for solving the hamming distance problem on exercism
package hamming

import (
	"errors"
)

//Distance Calculate the distance between two DNA strands
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return 0, errors.New("both strings must be the same length")
	}
	dist := 0
	for i := range ar {
		if ar[i] != br[i] {
			dist++
		}
	}
	return dist, nil
}
