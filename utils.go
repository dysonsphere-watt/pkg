package pkg

import "time"

const timeLayout = "15:04"
const dateLayout = "02-01-2006"

// Converts 24 hour time in the format "HH:mm" to a time.Time object
func StringToTime(timeStr string) (time.Time, error) {
	t, err := time.Parse(timeLayout, timeStr)
	return t, err
}

// Converts a time.Time object to a 24 hour string in the format "HH:mm"
func TimeToString(t time.Time) string {
	return t.Format(timeLayout)
}

// Converts a date string in the format "dd-MM-YYYY" to a time.Time object
func StringToDate(dateStr string) (time.Time, error) {
	t, err := time.Parse(dateLayout, dateStr)
	return t, err
}

// Converts a time.Time object to a date string in the format "dd-MM-YYYY"
func DateToString(t time.Time) string {
	return t.Format(dateLayout)
}