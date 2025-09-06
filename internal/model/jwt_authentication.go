package model

type JWTCheckResponse struct {
	Data struct {
		Email             string      `json:"email"`
		Levels            interface{} `json:"levels"`
		Nama              string      `json:"nama"`
		UserID            string      `json:"user_id"`
		Username          string      `json:"username"`
		IDPerguruanTinggi *int        `json:"id_perguruan_tinggi"`
	} `json:"data"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
