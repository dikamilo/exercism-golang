package clock

import "fmt"

const (
	minutesPerHour = 60
	hoursPerDay    = 24
)

type Clock struct {
	hour    int
	minutes int
}

func New(hour, minutes int) Clock {
	return Clock{hour, minutes}.Normalize()
}

func (c Clock) Add(minutes int) Clock {
	c.minutes += minutes
	return c.Normalize()
}

func (c Clock) Subtract(minutes int) Clock {
	c.minutes -= minutes
	return c.Normalize()
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minutes)
}

func (c Clock) Normalize() Clock {
	c.hour += c.minutes / minutesPerHour

	if c.minutes %= minutesPerHour; c.minutes < 0 {
		c.minutes += minutesPerHour
		c.hour--
	}

	if c.hour %= hoursPerDay; c.hour < 0 {
		c.hour += hoursPerDay
	}

	return c
}
