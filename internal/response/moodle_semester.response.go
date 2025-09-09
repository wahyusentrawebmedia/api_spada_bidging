package response

type MoodleSemesterResponse struct {
	Name        string `json:"name"`
	IDNumber    string `json:"idnumber" validate:"required"`
	Description string `json:"description"`
	Parent      string `json:"parent"`
}

type MoodleSemesterRequest struct {
	Name        string `json:"name" validate:"required"`
	IDNumber    string `json:"idnumber" validate:"required"`
	Description string `json:"description"`
	Parent      string `json:"parent"`
}
