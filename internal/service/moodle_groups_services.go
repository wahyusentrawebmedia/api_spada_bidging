package service

import (
	"api/spada/internal/model"
	"api/spada/internal/repository"
	"context"

	"gorm.io/gorm"
)

type MoodleGroupsService struct {
}

func NewMoodleGroupsService() *MoodleGroupsService {
	return &MoodleGroupsService{}
}

// GetGroupsByCategoriesID
func (s *MoodleGroupsService) GetGroupsByCategoriesID(idNumber string, db *gorm.DB) ([]*model.MdlGroups, error) {
	repo := repository.NewGroupsRepository(db)
	var repoCategories = repository.NewMoodleFakultasRepository(db)

	// Cek apakah Categories dengan IDNumber yang sama sudah ada
	existingCategories, err := repoCategories.GetFakultasByIDNumber(idNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	groups, err := repo.GetByCategoriesID(context.Background(), existingCategories.ID)
	if err != nil {
		return nil, err
	}
	return groups, nil
}
