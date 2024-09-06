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
	// Remove any enclosing quotes and extra spaces
	hhmmss = bytes.TrimSpace(bytes.Trim(hhmmss, `"`))

	if len(hhmmss) == 0 {
		return time.Time{}, fmt.Errorf("empty time input")
	}

	currentDate := time.Now().Format("2006-01-02")
	// Combine the current date with the provided time
	fullTimeStr := currentDate + " " + string(hhmmss)

	// Define the layouts corresponding to the possible time formats
	layouts := []string{
		"2006-01-02 15:04:05", // full time with seconds
		"2006-01-02 15:04",    // time without seconds
	}

	// Try parsing the time with each layout
	var parsedTime time.Time
	var err error
	for _, layout := range layouts {
		parsedTime, err = time.Parse(layout, fullTimeStr)
		if err == nil {
			return parsedTime, nil // return the first successful parse
		}
	}

	// If none of the layouts match, return the error
	return time.Time{}, fmt.Errorf("invalid time format '%s': %v", string(hhmmss), err)
}
