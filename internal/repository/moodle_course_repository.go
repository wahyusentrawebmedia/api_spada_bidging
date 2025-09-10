package repository

import (
	"api/spada/internal/model"
	"context"

	"gorm.io/gorm"
)

type MoodleCourseRepository struct {
	db *gorm.DB
}

func NewMoodleCourseRepository(db *gorm.DB) *MoodleCourseRepository {
	return &MoodleCourseRepository{db: db}
}

func (r *MoodleCourseRepository) Create(ctx context.Context, course *model.Course) error {
	return r.db.WithContext(ctx).Create(course).Error
}

func (r *MoodleCourseRepository) GetByID(ctx context.Context, id int64) (*model.Course, error) {
	var course model.Course
	if err := r.db.WithContext(ctx).First(&course, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &course, nil
}

// GetByIDNumber
func (r *MoodleCourseRepository) GetByIDNumber(idnumber string) (*model.Course, error) {
	var course model.Course
	if err := r.db.Where("idnumber = ?", idnumber).First(&course).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &course, nil
}

func (r *MoodleCourseRepository) Update(ctx context.Context, course *model.Course) error {
	return r.db.WithContext(ctx).Save(course).Error
}

func (r *MoodleCourseRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Course{}, id).Error
}

func (r *MoodleCourseRepository) List(ctx context.Context, limit, offset int) ([]*model.Course, error) {
	var courses []*model.Course
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}
