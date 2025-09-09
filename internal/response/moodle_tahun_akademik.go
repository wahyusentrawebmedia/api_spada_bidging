package response

type MoodleTahunAkademikResponse struct {
	Name        string `json:"name"`
	IDNumber    string `json:"idnumber" validate:"required"`
	Description string `json:"description"`
	Parent      string `json:"parent"`
}

type MoodleTahunAkademikRequest struct {
	Name        string `json:"name" validate:"required"`
	IDNumber    string `json:"idnumber" validate:"required"`
	Description string `json:"description"`
	Parent      string `json:"parent"`
}
