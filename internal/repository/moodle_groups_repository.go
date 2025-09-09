package repository

import (
	"api/spada/internal/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type GroupsRepository interface {
	Create(ctx context.Context, group *model.MdlGroups) error
	GetByID(ctx context.Context, id int64) (*model.MdlGroups, error)
	Update(ctx context.Context, group *model.MdlGroups) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*model.MdlGroups, error)
	GetByIDNumber(idnumber string) (*model.MdlGroups, error)
	GetByCategoriesID(ctx context.Context, categoriesID int64) ([]*model.MdlGroups, error)
}

type groupsRepository struct {
	db *gorm.DB
}

func NewGroupsRepository(db *gorm.DB) GroupsRepository {
	return &groupsRepository{db: db}
}

func (r *groupsRepository) Create(ctx context.Context, group *model.MdlGroups) error {
	if err := r.db.WithContext(ctx).Create(group).Error; err != nil {
		return err
	}
	return nil
}

// GetByIDNumber
func (r *groupsRepository) GetByIDNumber(idnumber string) (*model.MdlGroups, error) {
	var group model.MdlGroups
	if err := r.db.Where("idnumber = ?", idnumber).First(&group).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &group, nil
}

// GetByCategoriesID
func (r *groupsRepository) GetByCategoriesID(ctx context.Context, categoriesID int64) ([]*model.MdlGroups, error) {
	var groups []*model.MdlGroups
	if err := r.db.WithContext(ctx).Debug().Joins("JOIN mdl_course mc on mdl_groups.courseid = mc.id").
		Where("mc.category = ?", categoriesID).Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (r *groupsRepository) GetByID(ctx context.Context, id int64) (*model.MdlGroups, error) {
	var group model.MdlGroups
	if err := r.db.WithContext(ctx).First(&group, id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *groupsRepository) Update(ctx context.Context, group *model.MdlGroups) error {
	tx := r.db.WithContext(ctx).Debug().Model(&model.MdlGroups{}).Where("id = ?", group.ID).Updates(group)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no rows updated")
	}
	return nil
}

func (r *groupsRepository) Delete(ctx context.Context, id int64) error {
	tx := r.db.WithContext(ctx).Delete(&model.MdlGroups{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no rows deleted")
	}
	return nil
}

func (r *groupsRepository) List(ctx context.Context) ([]*model.MdlGroups, error) {
	var groups []*model.MdlGroups
	if err := r.db.WithContext(ctx).Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}
