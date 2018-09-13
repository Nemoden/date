package interval

import "time"

// We only represent interval by either months or seconds. Everything up to month can be represented as number of seconds,
// but onwards, month has a variadic length, and that's why a year may be represented only by number of months, but not number of days, for instance.
type Interval struct {
	Months  int
	Seconds int
}

func Months(months int) *Interval {
	return &Interval{Months: months}
}

func Years(years int) *Interval {
	return Months(12 * years)
}

func Seconds(seconds int) *Interval {
	return &Interval{Seconds: seconds}
}

func Minutes(minutes int) *Interval {
	return Seconds(minutes * 60)
}

func Hours(hours int) *Interval {
	return Minutes(hours * 60)
}

func Days(days int) *Interval {
	return Hours(days * 24)
}

func Weeks(weeks int) *Interval {
	return Days(weeks * 7)
}

// Adds interval to gived time.
func Add(t time.Time, i *Interval) time.Time {
	if i.Months != 0 {
		return t.AddDate(0, i.Months, 0)
	}
	if i.Seconds != 0 {
		return t.Add(time.Second * time.Duration(i.Seconds))
	}
	return t
}
