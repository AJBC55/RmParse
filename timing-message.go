package rmparse

import (
	"encoding/json"
	"errors"
	"fmt"
)

var MessageNotImplemented error = errors.New("message not Implemented ")

// struct for a timing message
type TimingMessage struct {
	Type string    `json:"type"`
	Data RmMessage `json:"data"`
}

func RmParse(line []byte) (*TimingMessage, error) {
	tm := new(TimingMessage)
	preped := prepLine(line)
	if preped == nil {
		return nil, fmt.Errorf("nil slice")
	}
	tm.Type = string(preped[0])
	switch string(preped[0]) {
	case "F":
		var rm Heartbeat
		if err := rm.RmParse(preped); err != nil {
			return nil, err
		}
		tm.Data = &rm
	case "COMP":
		var rm CompInfo
		if err := rm.RmParse(preped); err != nil {
			return nil, err
		}
		tm.Data = &rm
	case "B":
		var rm RunInfo
		if err := rm.RmParse(preped); err != nil {
			return nil, err
		}
		tm.Data = &rm
	case "C":
		var rm ClassInfo
		if err := rm.RmParse(preped); err != nil {
			return nil, err
		}
		tm.Data = &rm
	case "G":
		var rm RaceInfo
		if err := rm.RmParse(preped); err != nil {
			return nil, err
		}
		tm.Data = &rm
	case "H":
		var rm PracticeQualifyInfo
		if err := rm.RmParse(preped); err != nil {
			return nil, err
		}
		tm.Data = &rm
	case "J":
		var rm PassingInfo
		if err := rm.RmParse(preped); err != nil {
			return nil, err
		}
		tm.Data = &rm
	case "I":
		var rm InitRecord
		if err := rm.RmParse(preped); err != nil {
			return nil, err
		}
		tm.Data = &rm
	default:
		return nil, MessageNotImplemented

	}
	return tm, nil

}

// function to unmarshal tm json
func TmUnMarshal(dat []byte, tm *TimingMessage) error {
	temp := struct {
		Type string          `json:"type"`
		Data json.RawMessage `json:"data"`
	}{}
	if err := json.Unmarshal(dat, &temp); err != nil {
		return err
	}
	if temp.Type == "" || temp.Data == nil {
		return fmt.Errorf("jason err with temp")
	}
	tm.Type = temp.Type
	switch temp.Type {
	case "F":
		var rm Heartbeat
		if err := json.Unmarshal(temp.Data, &rm); err != nil {
			return err
		}
		tm.Data = &rm
	case "COMP":
		var rm CompInfo
		if err := json.Unmarshal(temp.Data, &rm); err != nil {
			return err
		}
		tm.Data = &rm
	case "B":
		var rm RunInfo
		if err := json.Unmarshal(temp.Data, &rm); err != nil {
			return err
		}
		tm.Data = &rm
	case "C":
		var rm ClassInfo
		if err := json.Unmarshal(temp.Data, &rm); err != nil {
			return err
		}
		tm.Data = &rm
	case "G":
		var rm RaceInfo
		if err := json.Unmarshal(temp.Data, &rm); err != nil {
			return err
		}
		tm.Data = &rm
	case "H":
		var rm PracticeQualifyInfo
		if err := json.Unmarshal(temp.Data, &rm); err != nil {
			return err
		}
		tm.Data = &rm
	case "J":
		var rm PassingInfo
		if err := json.Unmarshal(temp.Data, &rm); err != nil {
			return err
		}
		tm.Data = &rm
	case "I":
		var rm InitRecord
		if err := json.Unmarshal(temp.Data, &rm); err != nil {
			return err
		}
		tm.Data = &rm
	default:
		return MessageNotImplemented
	}
	return nil
}
