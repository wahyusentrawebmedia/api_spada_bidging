package repository

import (
	"api/spada/internal/model"
	"context"

	"gorm.io/gorm"
)

type MoodleGradesItemsRepository interface {
	Create(ctx context.Context, item *model.MdlGradeItems) error
	GetByID(ctx context.Context, id int64) (*model.MdlGradeItems, error)
	Update(ctx context.Context, item *model.MdlGradeItems) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]model.MdlGradeItems, error)
	FindByItemName(name *string, d *int64) (*model.MdlGradeItems, error)
}

type moodleGradesItemsRepository struct {
	db *gorm.DB
}

func NewMoodleGradesItemsRepository(db *gorm.DB) MoodleGradesItemsRepository {
	return &moodleGradesItemsRepository{db: db}
}

func (m *moodleGradesItemsRepository) FindByItemName(name *string, d *int64) (*model.MdlGradeItems, error) {
	var item model.MdlGradeItems
	err := m.db.Where("itemname = ? AND courseid = ?", name, d).First(&item).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

func (r *moodleGradesItemsRepository) Create(ctx context.Context, item *model.MdlGradeItems) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *moodleGradesItemsRepository) GetByID(ctx context.Context, id int64) (*model.MdlGradeItems, error) {
	var item model.MdlGradeItems
	err := r.db.WithContext(ctx).First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *moodleGradesItemsRepository) Update(ctx context.Context, item *model.MdlGradeItems) error {
	return r.db.WithContext(ctx).Save(item).Error
}

func (r *moodleGradesItemsRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.MdlGradeItems{}, id).Error
}

func (r *moodleGradesItemsRepository) List(ctx context.Context, limit, offset int) ([]model.MdlGradeItems, error) {
	var items []model.MdlGradeItems
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&items).Error
	return items, err
}
