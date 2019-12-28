package isogram

import (
	"strings"
	"unicode"
)

//IsIsogram determine whether the provided string is an isogram
func IsIsogram(s string) bool {
	if len(s) == 0 {
		return true
	}
	lowered := strings.ToLower(s)
	count := make(map[rune]int)
	for _, val := range lowered {
		if !unicode.IsLetter(val) {
			continue
		}
		_, ok := count[val]
		if ok {
			return false
		}
		count[val] = 1
	}
	return true
}
