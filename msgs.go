package rmparse

import (
	"bytes"
	"strconv"
	"time"
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

// struct for InitRecord
type InitRecord struct {
	TimeOfDay time.Time `json:"timeOfDay"`
	Date      time.Time `json:"date"`
}

// method for  Rm Format
func (ir *InitRecord) RmFormat(ln [][]byte) error {
	tm, err := parseTimeWithCurrentDate(ln[0])
	if err != nil {
		return err
	}
	ir.TimeOfDay = tm
	tm, err = parseTimeWithCurrentDate(ln[1])
	if err != nil {
		return err
	}
	ir.Date = tm
	return nil

}

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

// struct for the PracticeQualifyInfo
type PracticeQualifyInfo struct {
	Position           int    `json:"position"`
	RegistrationNumber string `json:"registrationNumber"`
	BestLap            int    `json:"bestLap"`
	BestLaptime        string `json:"bestLaptime"` // Custom type to handle duration
}

// method for PracticeQualifyInfo RmParse
func (pq *PracticeQualifyInfo) RmParse(ln [][]byte) error {
	i, err := strconv.Atoi(string(ln[0]))
	if err != nil {
		return err
	}
	pq.Position = i
	pq.RegistrationNumber = string(bytes.Trim(ln[1], `"`))
	i, err = strconv.Atoi(string(ln[2]))
	if err != nil {
		return err
	}
	pq.BestLap = i
	pq.BestLaptime = string(bytes.Trim(ln[3], `"`))
	return nil

}

// struct for the raceInfo
type RaceInfo struct {
	Position           int           `json:"position"`
	RegistrationNumber string        `json:"registrationNumber"`
	Laps               int           `json:"laps"`
	TotalTime          time.Duration `json:"totalTime"` // Custom type to handle duration
}

// method for RaceInfo to parse the timing Message
func (ri *RaceInfo) RmParse(ln [][]byte) error {
	i, err := strconv.Atoi(string(ln[0]))
	if err != nil {
		return err
	}
	ri.Position = i
	ri.RegistrationNumber = string(bytes.Trim(ln[1], `"`))
	i, err = strconv.Atoi(string(ln[2]))
	if err != nil {
		return err
	}
	ri.Laps = i
	dur, err := parseDuration(ln[3])
	if err != nil {
		return err
	}
	ri.TotalTime = dur
	return nil

}

// struct for the RunInfo
type RunInfo struct {
	UniqueNumber int    `json:"uniqueNumber"`
	Description  string `json:"description"`
}

// method for RunInfo RmParse
func (ri *RunInfo) RmParse(ln [][]byte) error {
	i, err := strconv.Atoi(string(ln[0]))
	if err != nil {
		return err
	}
	ri.UniqueNumber = i
	ri.Description = string(bytes.Trim(ln[1], `"`))
	return nil
}

// struct for the SettingInfo
type SettingInfo struct {
	Description string `json:"description"`
	Value       string `json:"value"`
}

// TODO RmParse For Setting Info

// struct for the CorrectedFinish information
type CorrectedFinish struct {
	RegistrationNumber string `json:"registrationNumber"` // max 8 characters
	Number             string `json:"number"`             // max 5 characters
	Laps               int    `json:"laps"`               // 0 - 99999
	TotalTime          string `json:"totalTime"`          // "HH:MM:SS.DDD", Custom type to handle duration
	CorrectionTime     string `json:"correctionTime"`     // "+/-HH:MM:SS.DDD", Custom type to handle duration
}

// TODO RmParse Method For CorrectedFinish
