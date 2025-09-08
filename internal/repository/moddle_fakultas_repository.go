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

// GetFakultasByKode retrieves a fakultas by its kode
func (r *MoodleFakultasRepository) GetFakultasByIDNumber(kode string) (*model.MdlCourseCategory, error) {
	var fakultas model.MdlCourseCategory

	if err := r.db.Where("idnumber = ?", kode).First(&fakultas).Error; err != nil {
		return nil, err
	}
	return &fakultas, nil
}

// UpdateFakultas updates an existing fakultas in the database
func (r *MoodleFakultasRepository) UpdateFakultas(fakultas *model.MdlCourseCategory) error {
	return r.db.Save(fakultas).Error
}
