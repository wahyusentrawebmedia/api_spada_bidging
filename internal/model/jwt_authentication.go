package model

type JWTCheckResponse struct {
	Data struct {
		IDPerguruanTinggi int    `json:"id_perguruan_tinggi"`
		IDTahapAkademik   string `json:"id_tahap_akademik"`
		IDTahunAkademik   string `json:"id_tahun_akademik"`
		Token             string `json:"token"`
	} `json:"data"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type JWTUserCheckResponse struct {
	Data struct {
		Email             string      `json:"email"`
		IDPerguruanTinggi int         `json:"id_perguruan_tinggi"`
		Levels            interface{} `json:"levels"`
		Nama              string      `json:"nama"`
		UserID            string      `json:"user_id"`
		Username          string      `json:"username"`
	} `json:"data"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
