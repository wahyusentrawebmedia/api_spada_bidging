package model

import "time"

type PostgresConfig struct {
	IDPerguruanTinggi string     `json:"id_perguruan_tinggi" gorm:"column:id_perguruan_tinggi"`
	Name              string     `json:"name" gorm:"column:name"`
	Host              string     `json:"host" gorm:"column:host"`
	Port              int        `json:"port" gorm:"column:port"`
	User              string     `json:"user" gorm:"column:user"`
	Password          string     `json:"password" gorm:"column:password"`
	DBName            string     `json:"dbname" gorm:"column:dbname"`
	SSLMode           string     `json:"sslmode" gorm:"column:sslmode"`
	CreatedAt         time.Time  `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt         time.Time  `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt         *time.Time `json:"deleted_at" gorm:"column:deleted_at;index"`
	CreatedBy         string     `json:"created_by" gorm:"column:created_by"`
	UpdatedBy         string     `json:"updated_by" gorm:"column:updated_by"`
	DeletedBy         string     `json:"deleted_by" gorm:"column:deleted_by"`
}

// TableName sets the table name for PostgresConfig struct
func (PostgresConfig) TableName() string {
	return "postgres_config"
}
