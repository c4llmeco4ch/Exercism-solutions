package luhn

//Valid determine whether a provided string is a valid
func Valid(s string) bool {
	arr, ok := convertToIntSlice(s)
	if len(arr) <= 1 || !ok {
		return false
	}
	return calcValue(arr)%10 == 0
}

func convertToIntSlice(s string) ([]int, bool) {
	arr := make([]int, len(s))
	pos := len(s) - 1
	everyOther := 1
	numSpaces := 0
	for pos >= 0 {
		if s[pos] == ' ' {
			pos--
			numSpaces++
			continue
		}
		if int(s[pos]) < '0' || int(s[pos]) > '9' {
			return arr, false
		}
		val := int(s[pos] - '0')
		if everyOther%2 == 0 {
			val *= 2
			if val > 9 {
				val -= 9
			}
		}
		arr[pos] = val
		everyOther = (everyOther + 1) % 2
		pos--
	}
	return arr, len(s)-numSpaces > 1
}

func calcValue(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += int(val)
	}
	return sum
}
