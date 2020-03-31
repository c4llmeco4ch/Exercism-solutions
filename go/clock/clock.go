package clock

import "fmt"

//Clock a clock struct providing the hours and minutes of a certain time
type Clock struct {
	Minutes int
}

//New Create a new clock provided an hour and minute value
func New(h int, m int) Clock {
	for m < 0 {
		h--
		m += 60
	}
	for h < 0 {
		h += 24
	}
	return Clock{(m + h*60) % 1440} // 1440 = minutes in a day
}

//String the stringer function for a referenced clock
func (c Clock) String() string {
	h := 0
	m := c.Minutes
	if m >= 60 {
		h += m / 60
	}
	return fmt.Sprintf("%02d:%02d", h%24, m%60)
}

//Add provided some minutes 'm', add that much time to the referenced clock
func (c Clock) Add(m int) Clock {
	return New(0, c.Minutes+m)
}

//Subtract provided some minutes 'm', subtract that much time to the referenced clock
func (c Clock) Subtract(m int) Clock {
	return New(0, c.Minutes-m)
}
