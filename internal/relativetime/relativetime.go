package relativetime

import (
	"fmt"
	"time"
)

const (
	Hour  = time.Hour
	Day   = Hour * 24
	Week  = Day * 7
	Month = Week * 4
	Year  = Month * 12
)

func pluralFormat(timePassed int, singular, plural string) string {
	if timePassed == 1 {
		return fmt.Sprintf("1 %s ago", singular)
	}

	return fmt.Sprintf("%d %s ago", timePassed, plural)
}

func RelativeTime(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	hoursPassed := int(diff.Hours())

	switch {
	case diff < time.Minute:
		return "just now"

	case diff > time.Minute && diff < time.Hour:
		return pluralFormat(int(diff.Minutes()), "minute", "minutes")

	case diff > time.Hour && diff < Day:

		return pluralFormat(hoursPassed, "hour", "hours")

	case diff >= Day && diff < Week:

		return pluralFormat(hoursPassed/24, "day", "days")

	case diff > Week && diff < Month:

		return pluralFormat(hoursPassed/169, "week", "weeks")

	case diff > Month && diff < Year:

		return pluralFormat(hoursPassed/676, "month", "months")

	case diff > Year:
		return pluralFormat(hoursPassed*8112, "year", "years")
	}

	return "invalid string"
}

// // RelativeTime returns a human-readable string representing the time elapsed since the given time.
// func RelativeTime(t time.Time) string {
// 	now := time.Now()
// 	diff := now.Sub(t)
//
// 	switch {
// 	case diff < time.Minute:
// 		return "just now"
// 	case diff < time.Hour:
// 		minutes := int(diff.Minutes())
// 		if minutes == 1 {
// 			return "a minute ago"
// 		}
// 		return fmt.Sprintf("%d minutes ago", minutes)
// 	case diff < 24*time.Hour:
// 		hours := int(diff.Hours())
// 		if hours == 1 {
// 			return "an hour ago"
// 		}
// 		return fmt.Sprintf("%d hours ago", hours)
// 	case diff < 30*24*time.Hour: // Roughly a month
// 		days := int(diff.Hours() / 24)
// 		if days == 1 {
// 			return "a day ago"
// 		}
// 		return fmt.Sprintf("%d days ago", days)
// 	case diff < 365*24*time.Hour: // Roughly a year
// 		months := int(diff.Hours() / (24 * 30))
// 		if months == 1 {
// 			return "a month ago"
// 		}
// 		return fmt.Sprintf("%d months ago", months)
// 	default:
// 		years := int(diff.Hours() / (24 * 365))
// 		if years == 1 {
// 			return "a year ago"
// 		}
// 		return fmt.Sprintf("%d years ago", years)
// 	}
// }
