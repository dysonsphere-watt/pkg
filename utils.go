package pkg

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
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

// Converts 24 hour time in the format "HH:mm" to a time.Time object.
// Set dt to true if the date is required like for MySQL.
// Accepts an IANA timezone name to offset the time.
func StringToTimeTz(timeStr string, dt bool, ianaTz string) (time.Time, error) {
	loc, err := time.LoadLocation(ianaTz)
	if err != nil {
		return time.Now(), err
	}

	if dt {
		timeStr = fmt.Sprintf(dateTimeFmt, timeStr)
		return time.ParseInLocation(dateTimeLayout, timeStr, loc)
	}

	return time.ParseInLocation(timeLayout, timeStr, loc)
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

// Converts a date string in the format "YYYY-MM-dd" to a time.Time object.
// Accepts an IANA timezone name to offset the time.
func StringToDateTz(dateStr string, ianaTz string) (time.Time, error) {
	loc, err := time.LoadLocation(ianaTz)
	if err != nil {
		return time.Now(), err
	}

	return time.ParseInLocation(dateLayout, dateStr, loc)
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

/*
Should be generally used for HTTP requests, to check if the bodies are empty.
Given a JSON serialisable struct, check if any fields are empty given the requiredFields.
Parents should be an empty string slice when you call it.
Eg. The JSON body provided:

	{
		"service_ids": [1, 2, 3],
		"robot": {
			"name": "Robot1",
			"station_id": 123,
		},
	}

If you want ALL fields to be provided, have requiredFields = []string{"service_ids", "robot.name", "robot.station_id"}.
Nested fields are identified by all previous parents joined by a "."
*/
func CheckEmptyFields(obj interface{}, parents, requiredFields []string) ([]string, error) {
	var val reflect.Value
	missingFields := []string{}

	val = reflect.ValueOf(obj).Elem()

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("struct expected, received a %s", val.Kind().String())
	}

	requiredFieldsLookup := make(map[string]struct{})
	for _, field := range requiredFields {
		requiredFieldsLookup[field] = struct{}{}
	}

	for i := 0; i < val.Type().NumField(); i++ {
		field := val.Type().Field(i)
		jsonTag := field.Tag.Get("json")

		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		f := val.Field(i)
		curTagSlice := make([]string, len(parents)+1)
		copy(curTagSlice, parents)
		curTagSlice = append(curTagSlice, jsonTag)
		curTag := strings.Join(curTagSlice, ".")

		if f.Kind() == reflect.Ptr && f.Elem().Kind() == reflect.Struct {
			recurMissingFields, err := CheckEmptyFields(f.Interface(), curTagSlice, requiredFields)
			if err != nil {
				return nil, err
			}
			missingFields = append(missingFields, recurMissingFields...)
		} else {
			if _, ok := requiredFieldsLookup[curTag]; ok {
				if isEmptyValue(f) {
					missingFields = append(missingFields, curTag)
				}
			}
		}
	}

	return missingFields, nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return false
}
