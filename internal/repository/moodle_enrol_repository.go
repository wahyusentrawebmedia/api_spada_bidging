package repository

import (
	"api/spada/internal/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type EnrolRepository interface {
	Create(ctx context.Context, enrol *model.Enrol) error
	GetByID(ctx context.Context, id uint) (*model.Enrol, error)
	Update(ctx context.Context, enrol *model.Enrol) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]*model.Enrol, error)
	GetByCourseIDAndEnrol(ctx context.Context, courseID int64, enrolMethod string) (*model.Enrol, error)
}

type enrolRepository struct {
	db *gorm.DB
}

func NewMoodleEnrolRepository(db *gorm.DB) EnrolRepository {
	return &enrolRepository{db: db}
}

func (r *enrolRepository) Create(ctx context.Context, enrol *model.Enrol) error {
	return r.db.WithContext(ctx).Create(enrol).Error
}

func (r *enrolRepository) GetByID(ctx context.Context, id uint) (*model.Enrol, error) {
	var enrol model.Enrol
	if err := r.db.WithContext(ctx).First(&enrol, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &enrol, nil
}

func (r *enrolRepository) Update(ctx context.Context, enrol *model.Enrol) error {
	return r.db.WithContext(ctx).Save(enrol).Error
}

func (r *enrolRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Enrol{}, id).Error
}

func (r *enrolRepository) List(ctx context.Context) ([]*model.Enrol, error) {
	var enrols []*model.Enrol
	if err := r.db.WithContext(ctx).Find(&enrols).Error; err != nil {
		return nil, err
	}
	return enrols, nil
}

// GetByCourseIDAndEnrol retrieves enrolments by course ID
func (r *enrolRepository) GetByCourseIDAndEnrol(ctx context.Context, courseID int64, enrolMethod string) (*model.Enrol, error) {
	var enrol model.Enrol
	if err := r.db.WithContext(ctx).Where("courseid = ? AND enrol = ?", courseID, enrolMethod).First(&enrol).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &enrol, nil
}
