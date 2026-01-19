package timeutil

import (
	"strconv"
	"strings"
)

// ParseTimeToSeconds converts a time string in the format "HH:MM:SS" to seconds.
func ParseTimeToSeconds(timeStr string) float64 {
	parts := strings.Split(timeStr, ":")
	if len(parts) != 3 {
		return 0
	}
	h, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	s, _ := strconv.ParseFloat(parts[2], 64)
	return float64(h*3600+m*60) + s
}

