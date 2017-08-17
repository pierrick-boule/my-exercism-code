// Package clock implement clock utility.
package clock

import "fmt"

const testVersion = 4

// Type Clock hold data structure of clock.
type Clock struct {
	hour, minute int
}

// New function return a new clock.
func New(hour, minute int) Clock {
	if minute < 0 {
		hour += minute/60 - 1
		minute %= 60
		minute += 60
	}
	if hour < 0 {
		hour %= 24
		hour += 24
	}
	return Clock{(minute/60 + hour) % 24, minute % 60}
}

// String function return the time representation has a string.
func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

// Add allow to add minute to a clock.
func (clock Clock) Add(minutes int) Clock {
	return New(clock.hour, clock.minute+minutes)
}
