package luhn

import (
	"strings"
	"unicode"
)

//Valid determine whether a provided string is a valid
func Valid(s string) bool {
	arr, ok := convertToIntSlice(strings.ReplaceAll(s, " ", ""))
	if len(arr) <= 1 || !ok {
		return false
	}
	for pos := len(arr) - 2; pos >= 0; pos -= 2 {
		num := arr[pos] * 2
		if num > 9 {
			num -= 9
		}
		arr[pos] = num
	}
	return calcValue(arr)%10 == 0
}

func convertToIntSlice(s string) ([]int, bool) {
	arr := make([]int, len(s))
	for i, r := range s {
		if !unicode.IsDigit(r) {
			return arr, false
		}
		arr[i] = int(r - '0')
	}
	return arr, true
}

func calcValue(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += int(val)
	}
	return sum
}
