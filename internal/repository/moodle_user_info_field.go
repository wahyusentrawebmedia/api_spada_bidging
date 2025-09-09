package repository

import (
	"api/spada/internal/model"
	"context"

	"gorm.io/gorm"
)

type MoodleUserInfoFieldRepository interface {
	Create(ctx context.Context, field *model.UserInfoField) error
	GetByID(ctx context.Context, id uint) (*model.UserInfoField, error)
	Update(ctx context.Context, field *model.UserInfoField) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]model.UserInfoField, error)
	GetByShortName(ctx context.Context, shortName string) (*model.UserInfoField, error)
}

type moodleUserInfoFieldRepository struct {
	db *gorm.DB
}

func NewMoodleUserInfoFieldRepository(db *gorm.DB) MoodleUserInfoFieldRepository {
	return &moodleUserInfoFieldRepository{db: db}
}

func (r *moodleUserInfoFieldRepository) Create(ctx context.Context, field *model.UserInfoField) error {
	return r.db.WithContext(ctx).Create(field).Error
}

func (r *moodleUserInfoFieldRepository) GetByID(ctx context.Context, id uint) (*model.UserInfoField, error) {
	var field model.UserInfoField
	if err := r.db.WithContext(ctx).First(&field, id).Error; err != nil {
		return nil, err
	}
	return &field, nil
}

func (r *moodleUserInfoFieldRepository) Update(ctx context.Context, field *model.UserInfoField) error {
	return r.db.WithContext(ctx).Save(field).Error
}

func (r *moodleUserInfoFieldRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.UserInfoField{}, id).Error
}

func (r *moodleUserInfoFieldRepository) List(ctx context.Context) ([]model.UserInfoField, error) {
	var fields []model.UserInfoField
	if err := r.db.WithContext(ctx).Find(&fields).Error; err != nil {
		return nil, err
	}
	return fields, nil
}

// GetByShortName retrieves a UserInfoField by its short name
func (r *moodleUserInfoFieldRepository) GetByShortName(ctx context.Context, shortName string) (*model.UserInfoField, error) {
	var field model.UserInfoField
	if err := r.db.WithContext(ctx).Where("shortname = ?", shortName).First(&field).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &field, nil
}
