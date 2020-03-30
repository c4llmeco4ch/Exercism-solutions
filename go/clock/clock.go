package clock

import "fmt"

//Clock a clock struct providing the hours and minutes of a certain time
type Clock struct {
	Hours   int
	Minutes int
}

//New Create a new clock provided an hour and minute value
func New(h int, m int) Clock {
	hours := h
	if m >= 60 {
		hours += m / 60
	}
	for m < 0 {
		m += 60
		hours--
	}
	for hours < 0 {
		hours += 24
	}
	return Clock{hours % 24, m % 60}
}

//String the stringer function for a referenced clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}

//Add provided some minutes 'm', add that much time to the referenced clock
func (c *Clock) Add(m int) Clock {
	return New(c.Hours, c.Minutes+m)
}

//Subtract provided some minutes 'm', subtract that much time to the referenced clock
func (c *Clock) Subtract(m int) Clock {
	return New(c.Hours, c.Minutes-m)
}
