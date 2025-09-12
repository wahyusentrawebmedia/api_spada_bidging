package repository

import (
	"api/spada/internal/model"
	"context"
	"time"

	"gorm.io/gorm"
)

type RoleCapabilityRepository interface {
	Create(ctx context.Context, rc *model.RoleCapability) error
	GetByID(ctx context.Context, id int64) (*model.RoleCapability, error)
	Update(ctx context.Context, rc *model.RoleCapability) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]model.RoleCapability, error)
	SetCapability(ctx context.Context, roleID int64, capability string, capabilities []SetCapabilityParams) error
}

type roleCapabilityRepository struct {
	db *gorm.DB
}

func NewRoleCapabilityRepository(db *gorm.DB) RoleCapabilityRepository {
	return &roleCapabilityRepository{db: db}
}

func (r *roleCapabilityRepository) Create(ctx context.Context, rc *model.RoleCapability) error {
	return r.db.WithContext(ctx).Create(rc).Error
}

func (r *roleCapabilityRepository) GetByID(ctx context.Context, id int64) (*model.RoleCapability, error) {
	var rc model.RoleCapability
	if err := r.db.WithContext(ctx).First(&rc, id).Error; err != nil {
		return nil, err
	}
	return &rc, nil
}

func (r *roleCapabilityRepository) Update(ctx context.Context, rc *model.RoleCapability) error {
	return r.db.WithContext(ctx).Save(rc).Error
}

func (r *roleCapabilityRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.RoleCapability{}, id).Error
}

func (r *roleCapabilityRepository) List(ctx context.Context, limit, offset int) ([]model.RoleCapability, error) {
	var rcs []model.RoleCapability
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&rcs).Error; err != nil {
		return nil, err
	}
	return rcs, nil
}

type SetCapabilityParams struct {
	Name    string
	Allowed bool
}

// SetCapability sets or updates a capability for a given role in a specific context.
func (r *roleCapabilityRepository) SetCapability(ctx context.Context, roleID int64, capability string, capabilities []SetCapabilityParams) error {
	for _, cap := range capabilities {
		var rc model.RoleCapability
		err := r.db.WithContext(ctx).Where("roleid = ? AND capability = ?", roleID, cap.Name).First(&rc).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		permission := 0
		if cap.Allowed {
			permission = 1
		}

		if err == gorm.ErrRecordNotFound {
			// Create new capability
			newRC := model.RoleCapability{
				RoleID:       roleID,
				Capability:   cap.Name,
				Permission:   int64(permission),
				TimeModified: time.Now().Unix(),
			}
			if err := r.db.WithContext(ctx).Create(&newRC).Error; err != nil {
				return err
			}
		} else {
			// Update existing capability
			rc.Permission = int64(permission)
			rc.TimeModified = time.Now().Unix()
			if err := r.db.WithContext(ctx).Save(&rc).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
