package luhn

import (
	"strings"
	"unicode"
)

//Valid determine whether a provided string is a valid
func Valid(s string) bool {
	/*For some reason, the below comment runs slower (224 ns/op)
	* than the used line 20 (214 ns/op).
	* However, the commented section runs 5_426_456 times and completes in 1.446s
	* While the left in section runs 5_104_358 times and completes in 1.334s.
	* Not sure how to parse this information.
	* I'm reading this as both solutions are faster AND slower than each other
	* Which makes no sense.
	 */

	/*spliced := strings.ReplaceAll(s, " ", "")
	if len(spliced) <= 1 {
		return false
	}
	arr, ok := convertToIntSlice(spliced)
	if !ok {
		return false
	}*/
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
