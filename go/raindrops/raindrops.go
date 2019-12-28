package raindrops

import (
	"strconv"
	"strings"
)

const (
	pli int = 3
	pla int = 5
	plo int = 7
)

//Convert Take a number and give back the raindrop sound of that number
func Convert(val int) string {
	pling := playSound(val, pli, "Pling")
	plang := playSound(val, pla, "Plang")
	plong := playSound(val, plo, "Plong")
	answer := strings.Join([]string{pling, plang, plong}, "")
	if len(answer) == 0 {
		return strconv.Itoa(val)
	}
	return answer
}

func playSound(num int, mod int, s string) string {
	if num%mod == 0 {
		return s
	}
	return ""
}
