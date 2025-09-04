package response

import "api/spada/internal/model"

type PostgresConfigRequest struct {
	IDPerguruanTinggi int    `json:"id_perguruan_tinggi" binding:"required"`
	Name              string `json:"name" binding:"required"`
	Host              string `json:"host" binding:"required"`
	Port              int    `json:"port" binding:"required"`
	User              string `json:"user" binding:"required"`
	Password          string `json:"password" binding:"required"`
	DBName            string `json:"dbname" binding:"required"`
	SSLMode           string `json:"sslmode"`
	Endpoint          string `json:"endpoint" binding:"required"`
}

func (req PostgresConfigRequest) ToModel() model.PostgresConfig {
	return model.PostgresConfig{
		IDPerguruanTinggi: req.IDPerguruanTinggi,
		Name:              req.Name,
		Host:              req.Host,
		Port:              req.Port,
		User:              req.User,
		Password:          req.Password,
		DBName:            req.DBName,
		SSLMode:           req.SSLMode,
		Endpoint:          req.Endpoint,
	}
}
