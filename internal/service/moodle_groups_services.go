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

	groups, err := repo.GetByCategoriesIDNumber(context.Background(), idNumber)
	if err != nil {
		return nil, err
	}
	return groups, nil
}
