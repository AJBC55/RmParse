package rmparse

import (
	"bytes"
	"strconv"
)

// struct for the ClassInfo
type ClassInfo struct {
	UniqueNumber int    `json:"uniqueNumber"`
	Description  string `json:"description"`
}

// method for rm parse
func (ci *ClassInfo) RmParse(ln [][]byte) error {
	un, err := strconv.Atoi(string(ln[0]))
	if err != nil {
		return err
	}
	ci.UniqueNumber = un
	ci.Description = string(bytes.Trim(ln[1], `"`))
	return nil
}

// struct for the CorrectedFinish
type CorrectedFinish struct {
	RegistrationNumber string `json:"registrationNumber"` // max 8 characters
	Number             string `json:"number"`             // max 5 characters
	Laps               int    `json:"laps"`               // 0 - 99999
	TotalTime          string `json:"totalTime"`          // "HH:MM:SS.DDD", Custom type to handle duration
	CorrectionTime     string `json:"correctionTime"`     // "+/-HH:MM:SS.DDD", Custom type to handle duration
}

// struct for the Comp Info
type CompInfo struct {
	RegistrationNumber string `json:"registrationNumber"`
	Number             string `json:"number"`
	ClassNumber        int    `json:"classNumber"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Nationality        string `json:"nationality"`
}

// rm parse Method
func (cmp *CompInfo) RmParse(ln [][]byte) error {
	cmp.RegistrationNumber = string(bytes.Trim(ln[0], `"`))
	cmp.Number = string(bytes.Trim(ln[1], `"`))
	cn, err := strconv.Atoi(string(ln[2]))
	if err != nil {
		return err
	}
	cmp.ClassNumber = cn
	cmp.FirstName = string(bytes.Trim(ln[3], `"`))
	cmp.LastName = string(bytes.Trim(ln[4], `"`))
	cmp.Nationality = string(bytes.Trim(ln[5], `"`))
	return nil
}
