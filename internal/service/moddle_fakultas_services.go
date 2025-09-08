package service

import (
	"api/spada/internal/model"
	"api/spada/internal/repository"
	"api/spada/internal/response"

	"gorm.io/gorm"
)

type MoodleFakultasService struct {
}

func NewMoodleFakultasService() *MoodleFakultasService {
	return &MoodleFakultasService{}
}

// AddFakultas adds a new fakultas to the database
func (s *MoodleFakultasService) AddFakultas(req response.MoodleFakultasRequest, db *gorm.DB) (*model.MdlCourseCategory, error) {
	var repoFakultas = repository.NewMoodleFakultasRepository(db)
	var repoContext = repository.NewMoodleContextRepository(db)

	var fakultas model.MdlCourseCategory

	fakultas.Name = req.Name
	fakultas.IDNumber = &req.IDNumber
	fakultas.Description = &req.Description

	if err := repoFakultas.AddNewFakultas(&fakultas); err != nil {
		return nil, err
	}

	context := model.MdlContext{
		ContextLevel: 40, // Level for course category
		InstanceID:   fakultas.ID,
	}

	if err := repoContext.Create(nil, &context); err != nil {
		return nil, err
	}

	return &fakultas, nil
}

// GetFakultas retrieves all fakultas from the database
func (s *MoodleFakultasService) GetFakultas(req response.MoodleFakultasRequest, db *gorm.DB) ([]model.MdlCourseCategory, error) {
	var repoFakultas = repository.NewMoodleFakultasRepository(db)

	var fakultas []model.MdlCourseCategory

	fakultas, err := repoFakultas.GetAllFakultas()
	if err != nil {
		return nil, err
	}

	return fakultas, nil
}
