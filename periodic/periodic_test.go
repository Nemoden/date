package periodic

import "testing"
import "time"
import "github.com/Nemoden/date/periodic"

type testPair struct {
	date     time.Time
	expected bool
}

type testPairs []testPair

func makeDate(date string) time.Time {
	t, e := time.Parse("2 Jan 2006", date)
	if e != nil {
		panic(e)
	}
	return t
}

func makeTestPair(date string, expected bool) testPair {
	return testPair{makeDate(date), expected}
}

func TestPeriodicMonths(t *testing.T) {
	start := makeDate("31 Dec 2000")
	end := makeDate("31 Dec 2020")
	monthly := periodic.Monthly(start, end)

	pairs := testPairs{
		makeTestPair("31 Jan 2001", true),
		makeTestPair("3 Mar 2001", true),
		makeTestPair("28 Feb 2001", false),
	}

	for _, pair := range pairs {
		if monthly.Matches(pair.date) != pair.expected {
			t.Error(pair.date, "should match", monthly.StartsAt)
		}
	}
}

func TestPeriodicDays(t *testing.T) {
	start := makeDate("31 Dec 2000")
	end := makeDate("31 Dec 2020")
	each4Days := periodic.OnceInXDays(start, end, 4)

	pairs := testPairs{
		makeTestPair("3 Jan 2001", false),
		makeTestPair("4 Jan 2001", true),
		makeTestPair("5 Jan 2001", false),
		makeTestPair("28 Jan 2001", true),
		makeTestPair("1 Feb 2001", true),
		makeTestPair("25 Feb 2001", true),
		// Test asbcense of 29 of Feb, so that 25 Feb + 4 days would be 1 or Mar.
		makeTestPair("1 Mar 2001", true),
	}

	for _, pair := range pairs {
		if each4Days.Matches(pair.date) != pair.expected {
			t.Error(pair.date, "should match", each4Days.StartsAt)
		}
	}
}
