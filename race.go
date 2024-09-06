package rmparse

type RaceInfo struct {
	Position           int    `json:"position"`
	RegistrationNumber string `json:"registrationNumber"`
	Laps               int    `json:"laps"`
	TotalTime          string `json:"totalTime"` // Custom type to handle duration
}
