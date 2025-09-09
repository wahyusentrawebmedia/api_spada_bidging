package response

type PostgreConfigResponse struct {
	IDPerguruanTinggi string `json:"id_perguruan_tinggi" gorm:"column:id_perguruan_tinggi"`
	Type              string `json:"type" gorm:"column:type;type:VARCHAR(10)"` // psql, mysql
	Host              string `json:"host" gorm:"column:host"`
	Port              int    `json:"port" gorm:"column:port"`
	User              string `json:"user" gorm:"column:user"`
	Password          string `json:"password" gorm:"column:password"`
	DBName            string `json:"dbname" gorm:"column:dbname"`
	SSLMode           string `json:"sslmode" gorm:"column:sslmode"`
}

type PostgreConfigRequest struct {
	IDPerguruanTinggi string `json:"id_perguruan_tinggi" gorm:"column:id_perguruan_tinggi"`
	Type              string `json:"type" gorm:"column:type;type:VARCHAR(10)"` // psql, mysql
	Host              string `json:"host" gorm:"column:host"`
	Port              int    `json:"port" gorm:"column:port"`
	User              string `json:"user" gorm:"column:user"`
	Password          string `json:"password" gorm:"column:password"`
	DBName            string `json:"dbname" gorm:"column:dbname"`
	SSLMode           string `json:"sslmode" gorm:"column:sslmode"`
}
