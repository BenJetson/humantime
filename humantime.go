package humantime

import (
	"fmt"
	"time"
)

const (
	// Note that hours per month is approximate due to fluctuations in month length.
	HOURS_IN_MONTH = 730
	HOURS_IN_DAY = 24
	HOURS_IN_WEEK = HOURS_IN_DAY * 7
	HOURS_IN_YEAR = HOURS_IN_DAY * 365
)

// Duration returns a string with a plain English descrption of the length of
// time that the time.Duration t contains.
func Duration(t time.Duration) string {
	hours := int(t.Hours())
	minutes := int(t.Minutes())
	seconds := int(t.Seconds())

	if hours >= HOURS_IN_YEAR {
		years := hours / HOURS_IN_YEAR

		if years == 1 {
			return "a year ago"
		}
		return fmt.Sprintf("%d years ago", years)
	}

	if hours >= HOURS_IN_MONTH {
		months := hours / HOURS_IN_MONTH

		if months == 1 {
			return "about a month ago"
		}
		return fmt.Sprintf("about %d months ago", months)
	}

	if hours >= HOURS_IN_WEEK {
		weeks := hours / HOURS_IN_WEEK

		if weeks == 1 {
			return "a week ago"
		}
		return fmt.Sprintf("%d weeks ago", weeks)
	}

	if hours >= HOURS_IN_DAY {
		days := hours / HOURS_IN_DAY

		if days == 1 {
			return "a day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	}

	if hours > 0 {
		if hours == 1 {
			return "an hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	}

	if minutes > 0 {
		if minutes == 1 {
			return "a minute ago"
		}
		return fmt.Sprintf("%d minutes ago", minutes)
	}

	if seconds > 20 {
		return "seconds ago"
	}

	return "just now"
}

// Since provides an easy way to call Duration without having to call time.Since
// on a time.Time first.
func Since(t time.Time) string {
	return Duration(time.Since(t))
}
