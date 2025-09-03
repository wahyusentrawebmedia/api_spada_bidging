package service

import (
	"api/spada/internal/model"

	"gorm.io/gorm"
)

type KategoriService struct {
	// Add dependencies here, e.g. repository
	db *gorm.DB
}

func NewKategoriService(db *gorm.DB) *KategoriService {
	return &KategoriService{db: db}
}

// GetAllKategori retrieves all categories from the database
func (s *KategoriService) GetAllKategori(req model.RequestKategori) ([]model.Kategori, error) {
	var kategoris []model.Kategori

	return kategoris, nil
}
