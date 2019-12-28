package diffsquares

import "math"

//Difference Given a number, calculate the difference between the sum of squares and the square of sum
func Difference(n int) int {
	return int(math.Abs(float64(SumOfSquares(n) - SquareOfSum(n))))
}

//SumOfSquares Given a number, return the sum of the squares from 1 to n
func SumOfSquares(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return (n * n) + SumOfSquares(n-1)
}

//SquareOfSum Given a number, calculate the sum from 1 to n, then return the square of that sum
func SquareOfSum(n int) int {
	sum := n
	if n%2 == 0 {
		sum += ((n - 1) * (n / 2))
	} else {
		sum += (n * (n / 2))
	}
	return sum * sum
}
