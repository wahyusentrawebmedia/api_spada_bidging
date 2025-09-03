package response

type PostgresConfigRequest struct {
	IDPerguruanTinggi string `json:"id_perguruan_tinggi"`
	Name              string `json:"name"`
	Host              string `json:"host"`
	Port              int    `json:"port"`
	User              string `json:"user"`
	Password          string `json:"password"`
	DBName            string `json:"dbname"`
	SSLMode           string `json:"sslmode"`
}
