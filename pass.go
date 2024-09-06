package rmparse

import (
	"bytes"
	"time"
)

// struct for the passing Info
type PassingInfo struct {
	RegistrationNumber string        `json:"registrationNumber"`
	LapTime            time.Duration `json:"lapTime"`   // Custom type to handle duration
	TotalTime          time.Duration `json:"totalTime"` // Custom type to handle duration
}

func (pi *PassingInfo) RmParse(ln [][]byte) error {
	pi.RegistrationNumber = string(bytes.Trim(ln[0], `"`))
	dur, err := parseDuration(ln[1])
	if err != nil {
		return err
	}
	pi.LapTime = dur
	dur, err = parseDuration(ln[2])
	if err != nil {
		return err
	}
	pi.TotalTime = dur
	return nil

}
