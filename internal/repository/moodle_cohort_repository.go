package repository

import (
	"api/spada/internal/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type MoodleCohortRepository interface {
	Create(ctx context.Context, cohort *model.Cohort) error
	GetByID(ctx context.Context, id uint) (*model.Cohort, error)
	Update(ctx context.Context, cohort *model.Cohort) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]*model.Cohort, error)
	GetCohortByIDNumber(idNumber string) (*model.Cohort, error)
}

type moodleCohortRepository struct {
	db *gorm.DB
}

func NewMoodleCohortRepository(db *gorm.DB) MoodleCohortRepository {
	return &moodleCohortRepository{db: db}
}

func (r *moodleCohortRepository) Create(ctx context.Context, cohort *model.Cohort) error {
	return r.db.WithContext(ctx).Create(cohort).Error
}

func (r *moodleCohortRepository) GetByID(ctx context.Context, id uint) (*model.Cohort, error) {
	var cohort model.Cohort
	if err := r.db.WithContext(ctx).First(&cohort, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &cohort, nil
}

func (r *moodleCohortRepository) Update(ctx context.Context, cohort *model.Cohort) error {
	return r.db.WithContext(ctx).Save(cohort).Error
}

func (r *moodleCohortRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Cohort{}, id).Error
}

func (r *moodleCohortRepository) List(ctx context.Context) ([]*model.Cohort, error) {
	var cohorts []*model.Cohort
	if err := r.db.WithContext(ctx).Find(&cohorts).Error; err != nil {
		return nil, err
	}
	return cohorts, nil
}

// GetCohortByIDNumber retrieves a cohort by its IDNumber
func (r *moodleCohortRepository) GetCohortByIDNumber(idNumber string) (*model.Cohort, error) {
	var cohort model.Cohort
	if err := r.db.Where("idnumber = ?", idNumber).First(&cohort).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &cohort, nil
}
