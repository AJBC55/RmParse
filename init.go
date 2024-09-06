package rmparse

import "time"

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
