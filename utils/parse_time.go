package utils

import (
	"log"
	"time"
)

// ParseTime - parses a formatted string and returns the time value it represents.
func ParseTime(layout, timeString string) time.Time {
	// Parse dateString to time value
	t, err := time.Parse(layout, timeString)
	if err != nil {
		log.Fatalf("Error: Can't parse '%s' to time: %v\n", timeString, err.Error())
	}
	return t
}
