package luhn

//Valid determine whether a provided string is a valid
func Valid(s string) bool {
	numSpaces := 0
	sum := 0
	for pos := len(s) - 1; pos >= 0; pos-- {
		if s[pos] == ' ' {
			numSpaces++
			continue
		}
		if s[pos] < '0' || s[pos] > '9' {
			return false
		}
		val := int(s[pos] - '0')
		if pos%2 == (len(s)-numSpaces)%2 {
			val *= 2
			if val > 9 {
				val -= 9
			}
		}
		sum += val
	}
	return sum%10 == 0 && len(s)-numSpaces > 1
}
