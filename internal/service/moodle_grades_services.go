package service

import (
	"api/spada/internal/model"
	"api/spada/internal/repository"

	"gorm.io/gorm"
)

func NewMoodleGradesService() *MoodleGradesService {
	return &MoodleGradesService{}
}

type MoodleGradesService struct {
}

func (m MoodleGradesService) CreateAspekNilaiSync(db *gorm.DB, items *[]model.MdlGradeItems) ([]model.MdlGradeItems, error) {
	var createdItems []model.MdlGradeItems
	for _, item := range *items {
		createdItem, err := m.CreateAspekNilai(db, &item)
		if err != nil {
			return nil, err
		}
		createdItems = append(createdItems, *createdItem)
	}
	return createdItems, nil
}

func (m MoodleGradesService) CreateAspekNilai(db *gorm.DB, items *model.MdlGradeItems) (*model.MdlGradeItems, error) {
	repoGradeItems := repository.NewMoodleGradesItemsRepository(db)

	existing, err := repoGradeItems.FindByItemName(items.ItemName, items.CourseID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if existing != nil {
		// Update existing record
		items.ID = existing.ID
		err = repoGradeItems.Update(nil, items)
		if err != nil {
			return nil, err
		}
		return items, nil
	}

	err = repoGradeItems.Create(nil, items)
	if err != nil {
		return nil, err
	}
	return items, nil
}
