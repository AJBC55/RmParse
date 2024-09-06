package rmparse

import (
	"bytes"
	"strconv"
	"time"
)

type Heartbeat struct {
	LapsToGo   int           `json:"lapsToGo"`
	TimeToGo   time.Duration `json:"timeToGo"` // Custom type to handle duration
	TimeOfDay  time.Time     `json:"timeOfDay"`
	RaceTime   time.Duration `json:"raceTime"` // Custom type to handle duration
	FlagStatus string        `json:"flagStatus"`
}

// function for the rm parse for the hearbeat message
func (hb *Heartbeat) RmParse(formatedLine [][]byte) error {
	laps, err := strconv.Atoi(string(formatedLine[0]))
	if err != nil {
		return err
	}
	ttg, err := parseDuration(formatedLine[1])
	if err != nil {
		return err
	}
	tod, err := parseTimeWithCurrentDate(formatedLine[2])
	if err != nil {
		return err
	}
	rt, err := parseDuration(formatedLine[3])
	if err != nil {
		return err
	}
	fs := bytes.Trim(formatedLine[4], `"`)
	hb.LapsToGo = laps
	hb.TimeToGo = ttg
	hb.TimeOfDay = tod
	hb.RaceTime = rt
	hb.FlagStatus = string(fs)
	return nil
}
