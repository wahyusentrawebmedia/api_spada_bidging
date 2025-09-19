package service

import (
	"api/spada/internal/model"
	"api/spada/internal/repository"
	"api/spada/internal/response"
	"fmt"

	"gorm.io/gorm"
)

type MoodleCategoriesService struct {
}

// ptrString returns a pointer to the given string.
func ptrString(s string) *string {
	return &s
}

func NewMoodleCategoriesService() *MoodleCategoriesService {
	return &MoodleCategoriesService{}
}

// CreateourseCategories adds a new Categories to the database
func (s *MoodleCategoriesService) CreateourseCategories(req response.MoodleCategoriesRequest, db *gorm.DB) (*model.MdlCourseCategory, error) {
	var repoCategories = repository.NewMoodleFakultasRepository(db)
	var repoContext = repository.NewMoodleContextRepository(db)
	var repoCohort = repository.NewMoodleCohortRepository(db)

	// Cek apakah Categories dengan IDNumber yang sama sudah ada
	existingCategories, err := repoCategories.GetCourseCategoriesByIDNumber(req.IDNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// get parent
	idParent := int64(0)

	if req.Parent != "" {
		fakultas, err := repoCategories.GetCourseCategoriesByIDNumber(req.Parent)
		if err != nil {
			return nil, err
		}
		idParent = fakultas.ID
	}

	var Categories model.MdlCourseCategory

	// Hitung kedalaman (Deepth) berdasarkan parent-child relationship
	Deepth := 1
	ParentPath := ""
	currentParent := existingCategories.Parent
	for currentParent != 0 {
		parentCategory, err := repoCategories.GetCourseCategoriesByID(currentParent)
		if err != nil {
			break
		}
		Deepth++
		currentParent = parentCategory.Parent
		ParentPath = parentCategory.Path + "/" + ParentPath
	}

	if existingCategories != nil && existingCategories.ID > 0 {
		// Jika ada, update data Categories
		existingCategories.Name = req.Name
		existingCategories.Description = &req.Description
		existingCategories.Depth = int64(Deepth)
		// Set path dengan ID dia sendiri
		existingCategories.Path = fmt.Sprintf("/%d", existingCategories.ID)

		if idParent != 0 {
			existingCategories.Parent = idParent
		}

		if err := repoCategories.UpdateCourseCategories(existingCategories); err != nil {
			return nil, err
		}
		Categories = *existingCategories
	} else {
		// Jika tidak ada, buat baru
		Categories.Name = req.Name
		Categories.IDNumber = &req.IDNumber
		Categories.Description = &req.Description
		existingCategories.Depth = int64(Deepth)

		if idParent != 0 {
			Categories.Parent = idParent
		}

		if err := repoCategories.CreateCourseCategories(&Categories); err != nil {
			return nil, err
		}

		// Set path dengan ID dia sendiri
		Categories.Path = fmt.Sprintf("/%d", Categories.ID)

		// Update path di database setelah create
		if err := repoCategories.UpdateCourseCategories(&Categories); err != nil {
			return nil, err
		}
	}

	// carikan id parent dan path parent dari context
	parentContext, err := repoContext.GetByInstanceIDAndLevel(nil, int(idParent), 40)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	var deepthContext int64 = 1
	var parentContextPath string
	if parentContext != nil {
		deepthContext = int64(parentContext.Depth) + 1
		parentContextPath = *parentContext.Path
	}

	// buatkan context untuk course categories jika belum ada
	existingContext, err := repoContext.GetByInstanceIDAndLevel(nil, int(Categories.ID), 40)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if existingContext == nil || existingContext.ID == 0 {
		context := model.MdlContext{
			ContextLevel: 40, // Level for course category
			InstanceID:   Categories.ID,
			Path:         &parentContextPath,
		}

		if err := repoContext.Create(nil, &context); err != nil {
			existingContext.Path = ptrString(parentContextPath + fmt.Sprintf("/%d", existingContext.ID))
		}
	}
	// Update path dan depth jika context sudah ada
	existingContext.Depth = int8(deepthContext)
	existingContext.Path = ptrString(parentContextPath + fmt.Sprintf("/%d", existingContext.ID))

	if err := repoContext.Update(nil, existingContext); err != nil {
		return nil, err
	}

	// Cek apakah cohort dengan IDNumber yang sama sudah ada, jika tidak ada maka buatkan
	existingCohort, err := repoCohort.GetCohortByIDNumber(req.IDNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if existingCohort == nil || existingCohort.ID == 0 {
		cohort := model.Cohort{
			Name:     req.Name,
			IDNumber: &req.IDNumber,
			// Tambahkan field lain jika diperlukan
		}
		if err := repoCohort.Create(nil, &cohort); err != nil {
			return nil, err
		}
	}

	if len(req.Children) > 0 {
		for _, child := range req.Children {
			child.Parent = req.IDNumber
			_, err := s.CreateourseCategories(child, db)
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
func (s *MoodleCategoriesService) BatchCategoriesSync(req []response.MoodleCategoriesRequest, db *gorm.DB) ([]model.MdlCourseCategory, []error) {
	var createdCategories []model.MdlCourseCategory
	var errs []error
	for _, config := range req {
		category, err := s.CreateourseCategories(config, db)
		if err != nil {
			errs = append(errs, err)
		} else {
			createdCategories = append(createdCategories, *category)
		}
	}
	if len(errs) > 0 {
		return createdCategories, errs
	}
	return createdCategories, nil
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
