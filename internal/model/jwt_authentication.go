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
