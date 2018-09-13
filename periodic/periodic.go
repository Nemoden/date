package periodic

import "github.com/Nemoden/date/interval"
import "time"

type Periodic struct {
	Interval *interval.Interval
	StartsAt time.Time
	EndsAt   time.Time
}

// Checks if periodic event matches certain time.
func (p Periodic) Matches(t time.Time) bool {
	// TODO think about memoization here.
	from := p.StartsAt
	for ; (from.Before(t) || from.Equal(t)) && !p.Overdue(t); from = interval.Add(from, p.Interval) {
		if from.Equal(t) {
			return true
		}
	}
	return false
}

// Checks if periodic event is overdue.
func (p Periodic) Overdue(due time.Time) bool {
	return !p.EndsAt.IsZero() && due.After(p.EndsAt)
}

// Creates a periodic event which only runs once.
func Once(t time.Time) Periodic {
	return Periodic{StartsAt: t, EndsAt: time.Time{}, Interval: nil}
}

// Creates a periodic event which applies to each 7 days.
func Weekly(s, e time.Time) Periodic {
	return Periodic{StartsAt: s, EndsAt: e, Interval: interval.Weeks(1)}
}

// Creates a periodic event which applies to each 14 days.
func Fortnightly(s, e time.Time) Periodic {
	return Periodic{StartsAt: s, EndsAt: e, Interval: interval.Weeks(2)}
}

// Creates a periodic event which applies to each month.
func Monthly(s, e time.Time) Periodic {
	return Periodic{StartsAt: s, EndsAt: e, Interval: interval.Months(1)}
}

// Creates a periodic event which applies to each year.
func Anually(s, e time.Time) Periodic {
	return Periodic{StartsAt: s, EndsAt: e, Interval: interval.Years(1)}
}

// Creates a periodic event which applies for each N days.
func OnceInXDays(s, e time.Time, d int) Periodic {
	return Periodic{StartsAt: s, EndsAt: e, Interval: interval.Days(d)}
}
