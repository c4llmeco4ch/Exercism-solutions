package luhn

//Valid determine whether a provided string is a valid
func Valid(s string) bool {
	arr, sum, ok := convertToIntSlice(s)
	if len(arr) <= 1 || !ok {
		return false
	}
	return sum%10 == 0
}

func convertToIntSlice(s string) ([]int, int, bool) {
	arr := make([]int, len(s))
	pos := len(s) - 1
	everyOther := 1
	numSpaces := 0
	sum := 0
	for pos >= 0 {
		if s[pos] == ' ' {
			pos--
			numSpaces++
			continue
		}
		if int(s[pos]) < '0' || int(s[pos]) > '9' {
			return arr, -1, false
		}
		val := int(s[pos] - '0')
		if everyOther%2 == 0 {
			val *= 2
			if val > 9 {
				val -= 9
			}
		}
		arr[pos] = val
		sum += val
		everyOther = (everyOther + 1) % 2
		pos--
	}
	return arr, sum, len(s)-numSpaces > 1
}
