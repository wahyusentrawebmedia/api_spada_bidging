package repository

import (
	"api/spada/internal/model"
	"context"

	"gorm.io/gorm"
)

type MoodleRoleRepository interface {
	GetByID(ctx context.Context, id int64) (*model.Role, error)
	GetAll(ctx context.Context) ([]model.Role, error)
	Create(ctx context.Context, role *model.Role) error
	Update(ctx context.Context, role *model.Role) error
	Delete(ctx context.Context, id int64) error
}

type moodleRoleRepository struct {
	db *gorm.DB
}

func NewMoodleRoleRepository(db *gorm.DB) MoodleRoleRepository {
	return &moodleRoleRepository{db: db}
}

func (r *moodleRoleRepository) GetByID(ctx context.Context, id int64) (*model.Role, error) {
	var role model.Role
	if err := r.db.WithContext(ctx).First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *moodleRoleRepository) GetAll(ctx context.Context) ([]model.Role, error) {
	var roles []model.Role
	if err := r.db.WithContext(ctx).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *moodleRoleRepository) Create(ctx context.Context, role *model.Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *moodleRoleRepository) Update(ctx context.Context, role *model.Role) error {
	return r.db.WithContext(ctx).Save(role).Error
}

func (r *moodleRoleRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Role{}, id).Error
}
