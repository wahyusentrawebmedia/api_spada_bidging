package repository

import (
	"api/spada/internal/model"
	"context"

	"gorm.io/gorm"
)

type MoodleScaleRepository interface {
	Create(ctx context.Context, scale *model.MdlScale) error
	GetByID(ctx context.Context, id int64) (*model.MdlScale, error)
	Update(ctx context.Context, scale *model.MdlScale) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*model.MdlScale, error)
}

type moodleScaleRepository struct {
	db *gorm.DB
}

func NewMoodleScaleRepository(db *gorm.DB) MoodleScaleRepository {
	return &moodleScaleRepository{db: db}
}

func (r *moodleScaleRepository) Create(ctx context.Context, scale *model.MdlScale) error {
	return r.db.WithContext(ctx).Create(scale).Error
}

func (r *moodleScaleRepository) GetByID(ctx context.Context, id int64) (*model.MdlScale, error) {
	var scale model.MdlScale
	if err := r.db.WithContext(ctx).First(&scale, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &scale, nil
}

func (r *moodleScaleRepository) Update(ctx context.Context, scale *model.MdlScale) error {
	return r.db.WithContext(ctx).Save(scale).Error
}

func (r *moodleScaleRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.MdlScale{}, id).Error
}

func (r *moodleScaleRepository) List(ctx context.Context) ([]*model.MdlScale, error) {
	var scales []*model.MdlScale
	if err := r.db.WithContext(ctx).Find(&scales).Error; err != nil {
		return nil, err
	}
	return scales, nil
}
