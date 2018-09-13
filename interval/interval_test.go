package interval

import "testing"
import "time"
import "github.com/Nemoden/date/interval"

type testPair struct {
	initialDate   time.Time
	addedInterval *interval.Interval
	expectedDate  time.Time
}

type testPairs []testPair

func makeTestPair(initialDate string, addedInterval *interval.Interval, expectedDate string) testPair {
	return testPair{makeTime(initialDate), addedInterval, makeTime(expectedDate)}
}

func makeTime(date string) time.Time {
	t, e := time.Parse("2 Jan 2006, 15:04:05", date)
	if e != nil {
		panic(e)
	}
	return t
}

func TestInterval(t *testing.T) {
	pairs := testPairs{
		makeTestPair("1 Jan 2000, 00:00:00", interval.Seconds(1), "1 Jan 2000, 00:00:01"),
		makeTestPair("1 Jan 2000, 00:00:00", interval.Seconds(3600), "1 Jan 2000, 01:00:00"),
		makeTestPair("1 Jan 2000, 00:00:00", interval.Seconds(86400), "2 Jan 2000, 00:00:00"),
		makeTestPair("1 Jan 2000, 00:00:00", interval.Minutes(1), "1 Jan 2000, 00:01:00"),
		makeTestPair("1 Jan 2000, 00:00:00", interval.Minutes(60), "1 Jan 2000, 01:00:00"),
		makeTestPair("1 Jan 2000, 00:00:00", interval.Days(7), "8 Jan 2000, 00:00:00"),
		makeTestPair("1 Jan 2000, 00:00:00", interval.Days(30), "31 Jan 2000, 00:00:00"),
		makeTestPair("1 Jan 2000, 00:00:00", interval.Days(60), "1 Mar 2000, 00:00:00"),
		makeTestPair("1 Jan 2000, 00:00:00", interval.Days(365), "31 Dec 2000, 00:00:00"),
		makeTestPair("1 Jan 2000, 00:00:00", interval.Days(366), "1 Jan 2001, 00:00:00"),
		makeTestPair("1 Jan 2000, 00:00:00", interval.Weeks(4), "29 Jan 2000, 00:00:00"),
		makeTestPair("1 Jan 2000, 00:00:00", interval.Months(24), "1 Jan 2002, 00:00:00"),
		makeTestPair("1 Jan 2000, 00:00:00", interval.Months(1), "1 Feb 2000, 00:00:00"),
		makeTestPair("1 Jan 2000, 00:00:00", interval.Months(2), "1 Mar 2000, 00:00:00"),
		makeTestPair("1 Jan 2000, 00:00:00", interval.Months(3), "1 Apr 2000, 00:00:00"),
		makeTestPair("1 Jan 2000, 00:00:00", interval.Years(1), "1 Jan 2001, 00:00:00"),
	}

	for _, pair := range pairs {
		actual := interval.Add(pair.initialDate, pair.addedInterval)
		if pair.expectedDate != actual {
			t.Error("Expected date:", pair.expectedDate, "actual date:", actual)
		}
	}
}
