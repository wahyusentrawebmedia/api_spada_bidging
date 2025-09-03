package repository

import "gorm.io/gorm"

func NewKategoriRepository(db *gorm.DB) *KategoriRepository {
	return &KategoriRepository{db: db}
}

type KategoriRepository struct {
	db *gorm.DB
}
