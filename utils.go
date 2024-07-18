package pkg

import (
	"encoding/json"
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

// Converts a date time string in the format "YYYY-MM-dd HH:mm" to a time.Time object
func StringToDateTime(dateTimeStr string) (time.Time, error) {
	t, err := time.Parse(dateTimeLayout, dateTimeStr)
	return t, err
}

// Converts a time.Time object to a datetime string in the format "YYYY-MM-dd HH:mm"
func DateTimeToString(t time.Time) string {
	return t.Format(dateTimeLayout)
}

// Converts a date time string in the format "YYYY-MM-dd HH:mm" to a time.Time object
// Accepts an IANA timezone name to offset the time.
func StringToDateTimeTz(dateTimeStr string, ianaTz string) (time.Time, error) {
	loc, err := time.LoadLocation(ianaTz)
	if err != nil {
		return time.Now(), err
	}

	t, err := time.ParseInLocation(dateTimeLayout, dateTimeStr, loc)
	return t, err
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

// Converts a struct with all string fields into a map
func StructToStrMap(obj interface{}) (map[string]string, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var result map[string]string
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Converts a struct into a JSON string
func StructToStr(obj interface{}) (string, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Ternary operator because go doesn't provide one because code cleanliness.
// DO NOT NEST unless you like dirty code.
func Ternary[T any](p bool, a, b T) T {
	if p {
		return a
	}
	return b
}
