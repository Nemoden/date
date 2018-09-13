Package `interval`
---

Shorthands to create interval:

    interval.Months(months int) *interval.Interval
    interval.Years(years int) *interval.Interval
    interval.Seconds(seconds int) *interval.Interval
    interval.Minutes(minutes int) *interval.Interval
    interval.Hours(hours int) *interval.Interval
    interval.Days(days int) *interval.Interval
    interval.Weeks(weeks int) *interval.Interval

Geenral usage:

    month := interval.Months(1)
    interval.Add(time.Now(), month) // Will correctly add 1 month (beware that months can not be represented as seconds as they are of variadic length)

    seconds := interval.Seconds(86400)
    interval.Add(time.Now(), seconds)

    // Absolutely the same:
    day := interval.Days(1)
    interval.Add(time.Now(), day)

Package `periodic`
---

`periodic.Periodic` simply represents periodic event (when it starts, when it ends and how often does it occur).

Geenral usage:

    start := time.Parse("2 Jan 2006", "1 Jan 2000")
    end := time.Parse("2 Jan 2006", "15 Jan 2000")
    fortnightly := periodic.Fortnightly(start, end) // fortnightly means once in 2 weeks ("bi-weekly").
    fortnightly.Matches(time.Parse("2 Jan 2006", "8 Jan 2000"))
    fortnightly.Matches(time.Parse("2 Jan 2006", "15 Jan 2000"))
    fortnightly.Matches(time.Parse("2 Jan 2006", "22 Jan 2000")) // false -> overdue!

You can omit periodic end date denoting that periodic event occurs indefinitely. This is done by specifying "zero" date as seconds param:

    start := time.Parse("2 Jan 2006", "1 Jan 2000")
    fortnightly := periodic.Fortnightly(start, time.Time{}) // runs indefinitely
    fortnightly.Matches(time.Parse("2 Jan 2006", "22 Jan 2000")) // true
