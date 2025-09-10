package repository

import (
	"api/spada/internal/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserEnrolmentRepository interface {
	Create(ctx context.Context, enrolment *model.UserEnrolment) error
	GetByID(ctx context.Context, id uint) (*model.UserEnrolment, error)
	Update(ctx context.Context, enrolment *model.UserEnrolment) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]*model.UserEnrolment, error)
	GetByEnrolIDAndUserID(ctx context.Context, enrolID int64, userID int64) (*model.UserEnrolment, error)
}

type userEnrolmentRepository struct {
	db *gorm.DB
}

func NewMoodleUserEnrolmentRepository(db *gorm.DB) UserEnrolmentRepository {
	return &userEnrolmentRepository{db: db}
}

func (r *userEnrolmentRepository) Create(ctx context.Context, enrolment *model.UserEnrolment) error {
	return r.db.WithContext(ctx).Create(enrolment).Error
}

func (r *userEnrolmentRepository) GetByID(ctx context.Context, id uint) (*model.UserEnrolment, error) {
	var enrolment model.UserEnrolment
	if err := r.db.WithContext(ctx).First(&enrolment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &enrolment, nil
}

func (r *userEnrolmentRepository) Update(ctx context.Context, enrolment *model.UserEnrolment) error {
	return r.db.WithContext(ctx).Save(enrolment).Error
}

func (r *userEnrolmentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.UserEnrolment{}, id).Error
}

func (r *userEnrolmentRepository) List(ctx context.Context) ([]*model.UserEnrolment, error) {
	var enrolments []*model.UserEnrolment
	if err := r.db.WithContext(ctx).Find(&enrolments).Error; err != nil {
		return nil, err
	}
	return enrolments, nil
}

// GetByEnrolIDAndUserID retrieves a user enrolment by enrol ID and user ID
func (r *userEnrolmentRepository) GetByEnrolIDAndUserID(ctx context.Context, enrolID int64, userID int64) (*model.UserEnrolment, error) {
	var enrolment model.UserEnrolment
	if err := r.db.WithContext(ctx).Where("enrolid = ? AND userid = ?", enrolID, userID).First(&enrolment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &enrolment, nil
}
