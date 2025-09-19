package service

import (
	"api/spada/internal/model"
	"api/spada/internal/repository"

	"gorm.io/gorm"
)

type MoodleGradesItemsService interface {
	CreateGradeItems(db *gorm.DB, req *model.GradeGrade) (*model.GradeGrade, error)
	CreateGradeItemsSync(db *gorm.DB, req *[]model.GradeGrade) (*[]model.GradeGrade, error)
}

func NewMoodleGradesGradesService() MoodleGradesItemsService {
	return &moodleGradesItemsService{}
}

type moodleGradesItemsService struct{}

func (s *moodleGradesItemsService) CreateGradeItems(db *gorm.DB, req *model.GradeGrade) (*model.GradeGrade, error) {
	repoGradesGrade := repository.NewMoodleGradesGradesRepository(db)

	existing, err := repoGradesGrade.GetByID(nil, req.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if existing != nil {
		// Update if exists
		if err := repoGradesGrade.Update(nil, req); err != nil {
			return nil, err
		}
		return req, nil
	}

	if err := repoGradesGrade.Create(nil, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (s *moodleGradesItemsService) CreateGradeItemsSync(db *gorm.DB, req *[]model.GradeGrade) (*[]model.GradeGrade, error) {
	var createdItems []model.GradeGrade
	for _, v := range *req {
		createdItem, err := s.CreateGradeItems(db, &v)
		if err != nil {
			return nil, err
		}
		createdItems = append(createdItems, *createdItem)
	}
	return &createdItems, nil
}
