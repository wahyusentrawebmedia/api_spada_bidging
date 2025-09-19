package repository

import (
	"api/spada/internal/model"
	"context"

	"gorm.io/gorm"
)

type MoodleGradesCategoriesRepository interface {
	Create(ctx context.Context, category *model.GradeCategory) error
	GetByID(ctx context.Context, id int64) (*model.GradeCategory, error)
	Update(ctx context.Context, category *model.GradeCategory) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, courseID int64) ([]model.GradeCategory, error)
}

type moodleGradesCategoriesRepository struct {
	db *gorm.DB
}

func NewMoodleGradesCategoriesRepository(db *gorm.DB) MoodleGradesCategoriesRepository {
	return &moodleGradesCategoriesRepository{db: db}
}

func (r *moodleGradesCategoriesRepository) Create(ctx context.Context, category *model.GradeCategory) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *moodleGradesCategoriesRepository) GetByID(ctx context.Context, id int64) (*model.GradeCategory, error) {
	var category model.GradeCategory
	err := r.db.WithContext(ctx).First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *moodleGradesCategoriesRepository) Update(ctx context.Context, category *model.GradeCategory) error {
	return r.db.WithContext(ctx).Save(category).Error
}

func (r *moodleGradesCategoriesRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.GradeCategory{}, id).Error
}

func (r *moodleGradesCategoriesRepository) List(ctx context.Context, courseID int64) ([]model.GradeCategory, error) {
	var categories []model.GradeCategory
	err := r.db.WithContext(ctx).Where("courseid = ?", courseID).Find(&categories).Error
	return categories, err
}
