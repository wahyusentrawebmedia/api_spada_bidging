package repository

import (
	"api/spada/internal/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type MdlCohortMemberRepository interface {
	Create(ctx context.Context, member *model.CohortMember) error
	GetByID(ctx context.Context, id int64) (*model.CohortMember, error)
	Update(ctx context.Context, member *model.CohortMember) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]model.CohortMember, error)
	GetByCohortIDAndUserId(ctx context.Context, cohortID int64, userID int64) (*model.CohortMember, error)
}

type mdlCohortMemberRepository struct {
	db *gorm.DB
}

func NewMoodleCohortMemberRepository(db *gorm.DB) MdlCohortMemberRepository {
	return &mdlCohortMemberRepository{db: db}
}

func (r *mdlCohortMemberRepository) Create(ctx context.Context, member *model.CohortMember) error {
	return r.db.WithContext(ctx).Create(member).Error
}

func (r *mdlCohortMemberRepository) GetByID(ctx context.Context, id int64) (*model.CohortMember, error) {
	var member model.CohortMember
	if err := r.db.WithContext(ctx).First(&member, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &member, nil
}

func (r *mdlCohortMemberRepository) Update(ctx context.Context, member *model.CohortMember) error {
	return r.db.WithContext(ctx).Save(member).Error
}

func (r *mdlCohortMemberRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.CohortMember{}, id).Error
}

func (r *mdlCohortMemberRepository) List(ctx context.Context) ([]model.CohortMember, error) {
	var members []model.CohortMember
	if err := r.db.WithContext(ctx).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

// GetByCohortIDAndUserId retrieves cohort members by cohort ID and user ID
func (r *mdlCohortMemberRepository) GetByCohortIDAndUserId(ctx context.Context, cohortID int64, userID int64) (*model.CohortMember, error) {
	var member model.CohortMember
	if err := r.db.WithContext(ctx).Where("cohortid = ? AND userid = ?", cohortID, userID).First(&member).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &member, nil
}
