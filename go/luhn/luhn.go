package luhn

//Valid determine whether a provided string is a valid
func Valid(s string) bool {
	realPos := 0
	sum := 0
	for pos := len(s) - 1; pos >= 0; pos-- {
		if s[pos] == ' ' {
			continue
		}
		if s[pos] < '0' || s[pos] > '9' {
			return false
		}
		val := int(s[pos] - '0')
		if realPos%2 == 1 {
			val *= 2
			if val > 9 {
				val -= 9
			}
		}
		realPos++
		sum += val
	}
	return sum%10 == 0 && realPos > 1
}
