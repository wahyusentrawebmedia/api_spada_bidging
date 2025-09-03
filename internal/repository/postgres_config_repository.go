package repository

import (
	"api/spada/internal/model"
	"context"

	"gorm.io/gorm"
)

type PostgresConfigRepository struct {
	db *gorm.DB
}

func NewPostgresConfigRepository(db *gorm.DB) *PostgresConfigRepository {
	return &PostgresConfigRepository{db: db}
}

func (r *PostgresConfigRepository) Create(ctx *context.Context, config *model.PostgresConfig) error {
	return r.db.Create(config).Error
}

func (r *PostgresConfigRepository) GetByID(ctx *context.Context, id int64) (*model.PostgresConfig, error) {
	var config model.PostgresConfig
	if err := r.db.First(&config, id).Error; err != nil {
		return nil, err
	}
	return &config, nil
}

func (r *PostgresConfigRepository) Update(ctx *context.Context, config *model.PostgresConfig) error {
	return r.db.Save(config).Error
}

func (r *PostgresConfigRepository) Delete(ctx *context.Context, id int64) error {
	return r.db.Delete(&model.PostgresConfig{}, id).Error
}

func (r *PostgresConfigRepository) List(ctx *context.Context) ([]*model.PostgresConfig, error) {
	var configs []*model.PostgresConfig
	if err := r.db.Find(&configs).Error; err != nil {
		return nil, err
	}
	return configs, nil
}
