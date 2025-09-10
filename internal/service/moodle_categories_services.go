package service

import (
	"api/spada/internal/model"
	"api/spada/internal/repository"
	"api/spada/internal/response"

	"gorm.io/gorm"
)

type MoodleCategoriesService struct {
}

func NewMoodleCategoriesService() *MoodleCategoriesService {
	return &MoodleCategoriesService{}
}

// AddCategories adds a new Categories to the database
func (s *MoodleCategoriesService) AddCategories(req response.MoodleCategoriesRequest, db *gorm.DB) (*model.MdlCourseCategory, error) {
	var repoCategories = repository.NewMoodleFakultasRepository(db)
	var repoContext = repository.NewMoodleContextRepository(db)

	// Cek apakah Categories dengan IDNumber yang sama sudah ada
	existingCategories, err := repoCategories.GetFakultasByIDNumber(req.IDNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// get parent
	idParent := int64(0)

	if req.Parent != "" {
		fakultas, err := repoCategories.GetFakultasByIDNumber(req.Parent)
		if err != nil {
			return nil, err
		}
		idParent = fakultas.ID
	}

	var Categories model.MdlCourseCategory

	if existingCategories != nil && existingCategories.ID > 0 {
		// Jika ada, update data Categories
		existingCategories.Name = req.Name
		existingCategories.Description = &req.Description

		if idParent != 0 {
			existingCategories.Parent = idParent
		}

		if err := repoCategories.UpdateFakultas(existingCategories); err != nil {
			return nil, err
		}
		Categories = *existingCategories
	} else {
		// Jika tidak ada, buat baru
		Categories.Name = req.Name
		Categories.IDNumber = &req.IDNumber
		Categories.Description = &req.Description

		if idParent != 0 {
			Categories.Parent = idParent
		}

		if err := repoCategories.AddNewFakultas(&Categories); err != nil {
			return nil, err
		}

		context := model.MdlContext{
			ContextLevel: 40, // Level for course category
			InstanceID:   Categories.ID,
		}

		if err := repoContext.Create(nil, &context); err != nil {
			return nil, err
		}
	}

	if len(req.Children) > 0 {
		for _, child := range req.Children {
			child.Parent = req.IDNumber
			_, err := s.AddCategories(child, db)
			if err != nil {
				return nil, err
			}
		}
	}

	return &Categories, nil
}

// GetCategories retrieves all Categories from the database
func (s *MoodleCategoriesService) GetCategories(db *gorm.DB) ([]model.MdlCourseCategory, error) {
	var repoCategories = repository.NewMoodleFakultasRepository(db)

	var Categories []model.MdlCourseCategory

	Categories, err := repoCategories.GetAllFakultas()
	if err != nil {
		return nil, err
	}

	return Categories, nil
}

// BatchCategoriesSync sync all Categories from all perguruan tinggi and returns a list of errors if any
func (s *MoodleCategoriesService) BatchCategoriesSync(req []response.MoodleCategoriesRequest, db *gorm.DB) []error {
	var errs []error
	for _, config := range req {
		_, err := s.AddCategories(config, db)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

// GetKeysWithPrefix returns all keys in the map that start with the given prefix
func (s *MoodleCategoriesService) GetCategoriesWithPrefix(prefix string, back bool, db *gorm.DB) ([]model.MdlCourseCategory, error) {
	var repoCategories = repository.NewMoodleFakultasRepository(db)

	var Categories []model.MdlCourseCategory

	if back {
		cat, err := repoCategories.GetWithPrefixEnd(prefix)
		if err != nil {
			return nil, err
		}
		Categories = append(Categories, cat...)
	} else {
		cat, err := repoCategories.GetWithPrefix(prefix)
		if err != nil {
			return nil, err
		}
		Categories = append(Categories, cat...)
	}

	return Categories, nil
}
