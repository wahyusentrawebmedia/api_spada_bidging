package repository

import (
	"api/spada/internal/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type RoleAssignmentRepository interface {
	Create(ctx context.Context, ra *model.RoleAssignment) error
	GetByID(ctx context.Context, id uint) (*model.RoleAssignment, error)
	Update(ctx context.Context, ra *model.RoleAssignment) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]model.RoleAssignment, error)
	GetByUserIDAndContextID(ctx context.Context, userID int64, contextID int64) (*model.RoleAssignment, error)
}

type roleAssignmentRepository struct {
	db *gorm.DB
}

func NewRoleAssignmentRepository(db *gorm.DB) RoleAssignmentRepository {
	return &roleAssignmentRepository{db: db}
}

func (r *roleAssignmentRepository) Create(ctx context.Context, ra *model.RoleAssignment) error {
	return r.db.WithContext(ctx).Create(ra).Error
}

func (r *roleAssignmentRepository) GetByID(ctx context.Context, id uint) (*model.RoleAssignment, error) {
	var ra model.RoleAssignment
	if err := r.db.WithContext(ctx).First(&ra, id).Error; err != nil {
		return nil, err
	}
	return &ra, nil
}

func (r *roleAssignmentRepository) Update(ctx context.Context, ra *model.RoleAssignment) error {
	return r.db.WithContext(ctx).Save(ra).Error
}

func (r *roleAssignmentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.RoleAssignment{}, id).Error
}

func (r *roleAssignmentRepository) List(ctx context.Context) ([]model.RoleAssignment, error) {
	var ras []model.RoleAssignment
	if err := r.db.WithContext(ctx).Find(&ras).Error; err != nil {
		return nil, err
	}
	return ras, nil
}

// GetByUserIDAndContextID retrieves a role assignment by user ID and context ID
func (r *roleAssignmentRepository) GetByUserIDAndContextID(ctx context.Context, userID int64, contextID int64) (*model.RoleAssignment, error) {
	var ra model.RoleAssignment
	if err := r.db.WithContext(ctx).Where("userid = ? AND contextid = ?", userID, contextID).First(&ra).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &ra, nil
}
