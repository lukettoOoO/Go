// go offers extensive support for times and durations
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// current time
	now := time.Now()
	p(now)

	// a time struct can be built by providing the year, month, day, etc.
	// times are always associated with a Location, i.e. a time zone
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// various components of the time value can be extracted as expected
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// the monday-sunday weekday is also available
	p(then.Weekday())

	// these methods compare two times, testing if the first occurs before, after, or at the same time as the second, respectively
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// the Sub methods returns a Duration representing the interval between two times
	diff := now.Sub(then)
	p(diff)

	// the length of the duration can be computed in various units
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Add can be used to advance a time by a given duration, or with a - to move backwards by a duration
	p(then.Add(diff))
	p(then.Add(-diff))
}
