package timeutil

import (
	"fmt"
	"time"
)

// FormatSecondsToTime converts a duration in seconds to a string in the format "HH:MM:SS".
func FormatSecondsToTime(seconds float64) string {
	h := int(seconds) / 3600
	m := (int(seconds) % 3600) / 60
	s := int(seconds) % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

// FormatDuration formats a time.Duration into a string in the format "HH:MM:SS".
func FormatDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
