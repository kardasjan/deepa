package utils

import (
	"strings"
	"time"

	"github.com/kardasjan/monday"
)

// ParseGeovisionDate date into time object
func ParseGeovisionDate(dateTime string) (time.Time, error) {
	parts := strings.SplitAfter(dateTime, ", ")
	loc, _ := time.LoadLocation("Europe/Prague")
	result, err := monday.ParseInLocation("02 Jan 06 15:04:05", parts[1], loc, monday.LocaleEnGB)
	if err != nil {
		return time.Now(), err
	}
	return result, nil
}
