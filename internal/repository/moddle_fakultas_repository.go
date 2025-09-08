package repository

import (
	"api/spada/internal/model"

	"gorm.io/gorm"
)

func NewMoodleFakultasRepository(db *gorm.DB) *MoodleFakultasRepository {
	return &MoodleFakultasRepository{db: db}
}

type MoodleFakultasRepository struct {
	db *gorm.DB
}

// AddNewFakultas adds a new fakultas to the database
func (r *MoodleFakultasRepository) AddNewFakultas(fakultas *model.MdlCourseCategory) error {
	return r.db.Create(fakultas).Error
}

// GetAllFakultas retrieves all fakultas from the database
func (r *MoodleFakultasRepository) GetAllFakultas() ([]model.MdlCourseCategory, error) {
	var fakultas []model.MdlCourseCategory

	if err := r.db.Debug().Find(&fakultas).Error; err != nil {
		return nil, err
	}
	return fakultas, nil
}
