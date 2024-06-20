package pkg

import (
	"fmt"
	"time"
)

const dateTimeFmt = "2010-01-01 %s"
const timeLayout = "15:04"
const dateLayout = "2006-01-02"
const dateTimeLayout = "2006-01-02 15:04"

// Converts 24 hour time in the format "HH:mm" to a time.Time object.
// Set dt to true if the date is required like for MySQL.
func StringToTime(timeStr string, dt bool) (time.Time, error) {
	if dt {
		timeStr = fmt.Sprintf(dateTimeFmt, timeStr)
		return time.Parse(dateTimeLayout, timeStr)
	}
	return time.Parse(timeLayout, timeStr)
}

// Converts a time.Time object to a 24 hour string in the format "HH:mm"
func TimeToString(t time.Time) string {
	return t.Format(timeLayout)
}

// Converts a date string in the format "YYYY-MM-dd" to a time.Time object
func StringToDate(dateStr string) (time.Time, error) {
	t, err := time.Parse(dateLayout, dateStr)
	return t, err
}

// Converts a time.Time object to a date string in the format "YYYY-MM-dd"
func DateToString(t time.Time) string {
	return t.Format(dateLayout)
}

// Converts a date time string in the format "YYYY-MM-dd HH:mm:ss" to a time.Time object
func StringToDateTime(dateTimeStr string) (time.Time, error) {
	t, err := time.Parse(dateTimeLayout, dateTimeStr)
	return t, err
}

// Converts a time.Time object to a datetime string in the format "YYYY-MM-dd HH:mm"
func DateTimeToString(t time.Time) string {
	return t.Format(dateTimeLayout)
}

// Takes a date and time, both being time.Time objects and combines the two.
// Accepts an IANA timezone name to offset the time.
func CombineDateAndTime(date, t time.Time, ianaTz string) time.Time {
	year, month, day := date.Date()
	hour, min, sec := t.Clock()

	loc, err := time.LoadLocation(ianaTz)
	if err != nil {
		loc = date.Location()
	}

	return time.Date(year, month, day, hour, min, sec, 0, loc)
}

// Ternary operator because go doesn't provide one because code cleanliness.
// DO NOT NEST unless you like dirty code.
func Ternary[T any](p bool, a, b T) T {
	if p {
		return a
	}
	return b
}
