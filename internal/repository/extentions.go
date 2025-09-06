package repository

import (
	"api/spada/internal/model"
	"api/spada/internal/utils"
	"fmt"

	"gorm.io/gorm"
)

func ToDBGorm(cfg *model.PostgresConfig) *gorm.DB {
	db, err := utils.ConnectionDB(cfg.User, cfg.Password, cfg.Host, fmt.Sprintf("%d", cfg.Port), cfg.DBName)
	if err != nil {
		return nil
	}
	return db
}
