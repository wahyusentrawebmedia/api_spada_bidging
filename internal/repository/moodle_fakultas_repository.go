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

// CreateCourseCategories adds a new fakultas to the database
func (r *MoodleFakultasRepository) CreateCourseCategories(fakultas *model.MdlCourseCategory) error {
	return r.db.Create(fakultas).Error
}

// GetAllFakultas retrieves all fakultas from the database
func (r *MoodleFakultasRepository) GetAllFakultas() ([]model.MdlCourseCategory, error) {
	var fakultas []model.MdlCourseCategory

	if err := r.db.Find(&fakultas).Error; err != nil {
		return nil, err
	}
	return fakultas, nil
}

// GetCourseCategoriesByID retrieves a fakultas by its ID
func (r *MoodleFakultasRepository) GetCourseCategoriesByID(id int64) (*model.MdlCourseCategory, error) {
	var fakultas model.MdlCourseCategory

	if err := r.db.First(&fakultas, id).Error; err != nil {
		return nil, err
	}
	return &fakultas, nil
}

// GetFakultasByKode retrieves a fakultas by its kode
func (r *MoodleFakultasRepository) GetCourseCategoriesByIDNumber(kode string) (*model.MdlCourseCategory, error) {
	var fakultas model.MdlCourseCategory

	if err := r.db.Where("idnumber = ?", kode).First(&fakultas).Error; err != nil {
		return nil, err
	}
	return &fakultas, nil
}

// UpdateCourseCategories updates an existing fakultas in the database
func (r *MoodleFakultasRepository) UpdateCourseCategories(fakultas *model.MdlCourseCategory) error {
	return r.db.Save(fakultas).Error
}

// GetAllProdi retrieves all prodi from the database
func (r *MoodleFakultasRepository) GetAllProdi(kodeFakultas string) ([]model.MdlCourseCategory, error) {
	var prodi []model.MdlCourseCategory

	var prodiFakultas model.MdlCourseCategory
	if err := r.db.Where("idnumber = ?", kodeFakultas).First(&prodiFakultas).Error; err != nil {
		return nil, err
	}

	if err := r.db.Where("parent = ?", prodiFakultas.ID).Find(&prodi).Error; err != nil {
		return nil, err
	}
	return prodi, nil
}

// GetWithPrefix retrieves all fakultas with a specific prefix in their IDNumber
func (r *MoodleFakultasRepository) GetWithPrefix(prefix string) ([]model.MdlCourseCategory, error) {
	var fakultas []model.MdlCourseCategory

	if err := r.db.Where("idnumber LIKE ?", prefix+"%").Find(&fakultas).Error; err != nil {
		return nil, err
	}
	return fakultas, nil
}

// GetWithPrefixEnd retrieves all fakultas with a specific prefix in their IDNumber
func (r *MoodleFakultasRepository) GetWithPrefixEnd(prefix string) ([]model.MdlCourseCategory, error) {
	var fakultas []model.MdlCourseCategory

	if err := r.db.Where("idnumber LIKE ?", "%"+prefix).Find(&fakultas).Error; err != nil {
		return nil, err
	}
	return fakultas, nil
}
