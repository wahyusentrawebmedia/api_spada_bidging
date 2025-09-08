package response

type MoodleFakultasResponse struct {
	Name        string `json:"name"`
	IDNumber    string `json:"idnumber"`
	Description string `json:"description"`
	// Parent      int    `json:"parent"`
}

type MoodleFakultasRequest struct {
	Name        string `json:"name" validate:"required"`
	IDNumber    string `json:"idnumber"`
	Description string `json:"description"`
	// Parent      int    `json:"parent"`
}
