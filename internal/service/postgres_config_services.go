package service

import (
	"api/spada/internal/model"
	"api/spada/internal/repository"
	"context"
)

type PostgresConfigService struct {
	repo repository.PostgresConfigRepository
}

func NewPostgresConfigService(repo repository.PostgresConfigRepository) *PostgresConfigService {
	return &PostgresConfigService{repo: repo}
}

func (s *PostgresConfigService) Create(ctx *context.Context, cfg *model.PostgresConfig) error {
	return s.repo.Create(ctx, cfg)
}

func (s *PostgresConfigService) GetByID(ctx *context.Context, id int64) (*model.PostgresConfig, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *PostgresConfigService) Update(ctx *context.Context, cfg *model.PostgresConfig) error {
	return s.repo.Update(ctx, cfg)
}

func (s *PostgresConfigService) Delete(ctx *context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *PostgresConfigService) List(ctx *context.Context) ([]*model.PostgresConfig, error) {
	return s.repo.List(ctx)
}
