package rmparse

import (
	"bytes"
	"fmt"
	"strconv"
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
	// Get the current date
	hhmmss = bytes.Trim(hhmmss, `"`)
	currentDate := time.Now().Format("2006-01-02")
	// Combine the current date with the provided time
	fullTimeStr := currentDate + " " + string(hhmmss)
	// Define the layout corresponding to the full date-time string
	layout := "2006-01-02 15:04:05"
	// Parse the combined date-time string to a time.Time object
	parsedTime, err := time.Parse(layout, fullTimeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time format: %v", err)
	}
	return parsedTime, nil
}
