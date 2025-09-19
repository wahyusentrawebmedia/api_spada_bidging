package repository

import (
	"api/spada/internal/model"
	"context"

	"gorm.io/gorm"
)

type MoodleGradesGradesRepository interface {
	Create(ctx context.Context, grade *model.GradeGrade) error
	GetByID(ctx context.Context, id int64) (*model.GradeGrade, error)
	Update(ctx context.Context, grade *model.GradeGrade) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]model.GradeGrade, error)
	GetByUserIDAndItemID(ctx context.Context, userID, itemID int64) (*model.GradeGrade, error)
}

type moodleGradesGradesRepository struct {
	db *gorm.DB
}

func NewMoodleGradesGradesRepository(db *gorm.DB) MoodleGradesGradesRepository {
	return &moodleGradesGradesRepository{db: db}
}

func (r *moodleGradesGradesRepository) GetByUserIDAndItemID(ctx context.Context, userID, itemID int64) (*model.GradeGrade, error) {
	var grade model.GradeGrade
	err := r.db.WithContext(ctx).Debug().
		Where("userid = ? AND itemid = ?", userID, itemID).
		First(&grade).Error
	if err != nil {
		return nil, err
	}
	return &grade, nil
}

func (r *moodleGradesGradesRepository) Create(ctx context.Context, grade *model.GradeGrade) error {
	return r.db.WithContext(ctx).Create(grade).Error
}

func (r *moodleGradesGradesRepository) GetByID(ctx context.Context, id int64) (*model.GradeGrade, error) {
	var grade model.GradeGrade
	err := r.db.WithContext(ctx).First(&grade, id).Error
	if err != nil {
		return nil, err
	}
	return &grade, nil
}

func (r *moodleGradesGradesRepository) Update(ctx context.Context, grade *model.GradeGrade) error {
	return r.db.WithContext(ctx).Save(grade).Error
}

func (r *moodleGradesGradesRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.GradeGrade{}, id).Error
}

func (r *moodleGradesGradesRepository) List(ctx context.Context, limit, offset int) ([]model.GradeGrade, error) {
	var grades []model.GradeGrade
	err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&grades).Error
	return grades, err
}
