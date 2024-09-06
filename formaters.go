package rmparse

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// function to prepare the rm message for parsing
func prepLine(line []byte) [][]byte {
	return commaSplit(cutFixes(line))
}

// function to split the trimed rm message
func commaSplit(line []byte) [][]byte {
	return bytes.Split(line, []byte(","))
}

// function for cuting the prefixes and suffixes of a Rmprotocal message
func cutFixes(line []byte) []byte {
	line = bytes.TrimPrefix(line, []byte("$"))
	line = bytes.TrimRight(line, "\n\r")
	return line
}

// function to turn the rm moniter duration format into a time duration
func parseDuration(hhmmss []byte) (time.Duration, error) {
	hhmmss = bytes.Trim(hhmmss, `"`)
	parts := bytes.Split(hhmmss, []byte(":"))
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid duration format")
	}
	hours, err := strconv.Atoi(string(parts[0]))
	if err != nil {
		return 0, fmt.Errorf("invalid hours: %v", err)
	}
	minutes, err := strconv.Atoi(string(parts[1]))
	if err != nil {
		return 0, fmt.Errorf("invalid minutes: %v", err)
	}
	seconds, err := strconv.ParseFloat(string(parts[2]), 64)
	if err != nil {
		return 0, fmt.Errorf("invalid seconds: %v", err)
	}
	duration := time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds*float64(time.Second))
	return duration, nil
}

// fucntion to turn the rmmoniter time format into a golang time
func parseTimeWithCurrentDate(hhmmss []byte) (time.Time, error) {
	// Remove any enclosing quotes and extra spaces
	hhmmss = bytes.TrimSpace(bytes.Trim(hhmmss, `"`))

	if len(hhmmss) == 0 {
		return time.Time{}, fmt.Errorf("empty time input")
	}

	// Split the time string by colon
	timeParts := strings.Split(string(hhmmss), ":")
	if len(timeParts) < 2 || len(timeParts) > 3 {
		return time.Time{}, fmt.Errorf("invalid time format, expected HH:mm or HH:mm:ss but got '%s'", string(hhmmss))
	}

	// Ensure hours, minutes, and seconds are always two digits
	if len(timeParts[0]) == 1 {
		timeParts[0] = "0" + timeParts[0] // Pad single digit hour
	}
	if len(timeParts[1]) == 1 {
		timeParts[1] = "0" + timeParts[1] // Pad single digit minute
	}

	if len(timeParts) == 2 {
		// If seconds part is missing, add "00"
		hhmmss = []byte(fmt.Sprintf("%s:%s:00", timeParts[0], timeParts[1]))
	} else if len(timeParts[2]) == 1 {
		// If the seconds part is a single digit, pad it with a leading zero
		hhmmss = []byte(fmt.Sprintf("%s:%s:%02s", timeParts[0], timeParts[1], timeParts[2]))
	} else {
		hhmmss = []byte(fmt.Sprintf("%s:%s:%s", timeParts[0], timeParts[1], timeParts[2]))
	}

	// Get the current date in the format "YYYY-MM-DD"
	currentDate := time.Now().Format("2006-01-02")

	// Combine the current date with the provided time
	fullTimeStr := currentDate + " " + string(hhmmss)

	// Define the layout corresponding to the full date-time string
	layout := "2006-01-02 15:04:05"

	// Parse the combined date-time string to a time.Time object
	parsedTime, err := time.Parse(layout, fullTimeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time format '%s': %v", string(hhmmss), err)
	}

	return parsedTime, nil
}
