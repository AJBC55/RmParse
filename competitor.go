package rmparse

import (
	"bytes"
	"strconv"
)

// struct for teh CompetitorInfo
type CompetitorInfo struct {
	RegistrationNumber string `json:"registrationNumber"`
	Number             string `json:"number"`
	TransponderNumber  int    `json:"transponderNumber"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Nationality        string `json:"nationality"`
	ClassNumber        int    `json:"classNumber"`
}

// rm Parse Function
func (ci *CompetitorInfo) RmParse(fl [][]byte) error {
	ci.RegistrationNumber = string(bytes.Trim(fl[0], `"`))
	ci.Number = string(bytes.Trim(fl[1], `"`))
	tn, err := strconv.Atoi(string(fl[2]))
	if err != nil {
		return err
	}
	ci.TransponderNumber = tn
	ci.FirstName = string(bytes.Trim(fl[3], `"`))
	ci.LastName = string(bytes.Trim(fl[4], `"`))
	ci.Nationality = string(bytes.Trim(fl[5], `"`))
	tn, err = strconv.Atoi(string(fl[2]))
	if err != nil {
		return err
	}
	ci.ClassNumber = tn
	return nil
}
