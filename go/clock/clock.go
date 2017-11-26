package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hours, minutes int
}

func New(hours, minutes int) Clock {
	var hrs, mins int
	time := hours*60 + minutes

	if time > 0 {
		for time >= 1440 {
			time = time - 1440
		}

		hrs = time / 60
		mins = time % 60

		return Clock{hours: hrs, minutes: mins}
	} else if time < 0 {
		for time < 0 {
			time = time + 1440
		}
		hrs = time / 60
		mins = time % 60
		return Clock{hours: hrs, minutes: mins}
	} else {
		return Clock{hours: 0, minutes: 0}
	}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.hours, c.minutes+minutes)
}
