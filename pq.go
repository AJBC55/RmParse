package rmparse

type PracticeQualifyInfo struct {
	Position           int    `json:"position"`
	RegistrationNumber string `json:"registrationNumber"`
	BestLap            int    `json:"bestLap"`
	BestLaptime        string `json:"bestLaptime"` // Custom type to handle duration
}
