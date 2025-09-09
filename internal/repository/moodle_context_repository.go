package repository

import (
	"api/spada/internal/model"
	"context"

	"gorm.io/gorm"
)

// MdlContextRepository defines CRUD operations for MdlContext.
type MdlContextRepository interface {
	Create(ctx context.Context, mdl *model.MdlContext) error
	GetByID(ctx context.Context, id int) (*model.MdlContext, error)
	Update(ctx context.Context, mdl *model.MdlContext) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]*model.MdlContext, error)
}

// MoodleContextRepository is an in-memory implementation of MdlContextRepository.
type MoodleContextRepository struct {
	db *gorm.DB
}

// NewMoodleContextRepository creates a new in-memory repository.
func NewMoodleContextRepository(db *gorm.DB) *MoodleContextRepository {
	return &MoodleContextRepository{
		db: db,
	}
}

func (r *MoodleContextRepository) Create(ctx context.Context, mdl *model.MdlContext) error {
	return r.db.WithContext(ctx).Create(mdl).Error
}

func (r *MoodleContextRepository) GetByID(ctx context.Context, id int) (*model.MdlContext, error) {
	var mdl model.MdlContext
	if err := r.db.WithContext(ctx).First(&mdl, id).Error; err != nil {
		return nil, err
	}
	return &mdl, nil
}

func (r *MoodleContextRepository) Update(ctx context.Context, mdl *model.MdlContext) error {
	return r.db.WithContext(ctx).Save(mdl).Error
}

func (r *MoodleContextRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&model.MdlContext{}, id).Error
}

func (r *MoodleContextRepository) List(ctx context.Context) ([]*model.MdlContext, error) {
	var contexts []*model.MdlContext
	if err := r.db.WithContext(ctx).Find(&contexts).Error; err != nil {
		return nil, err
	}
	return contexts, nil
}
