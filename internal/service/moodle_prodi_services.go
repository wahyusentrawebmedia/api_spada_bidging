package service

import (
	"api/spada/internal/model"
	"api/spada/internal/repository"
	"api/spada/internal/response"
	"errors"

	"gorm.io/gorm"
)

type MoodleProdiService struct {
}

func NewMoodleProdiService() *MoodleProdiService {
	return &MoodleProdiService{}
}

// AddProdi adds a new Prodi to the database
func (s *MoodleProdiService) AddProdi(req response.MoodleProdiRequest, db *gorm.DB) (*model.MdlCourseCategory, error) {
	var repoProdi = repository.NewMoodleFakultasRepository(db)
	var repoContext = repository.NewMoodleContextRepository(db)

	if req.Parent == "" {
		return nil, errors.New("parent is required")
	}

	fakultas, err := repoProdi.GetFakultasByIDNumber(req.Parent)
	if err != nil {
		return nil, err
	}

	// Cek apakah Prodi dengan IDNumber yang sama sudah ada
	existingProdi, err := repoProdi.GetFakultasByIDNumber(req.IDNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	var Prodi model.MdlCourseCategory

	if existingProdi != nil && existingProdi.ID > 0 {
		// Jika ada, update data Prodi
		existingProdi.Name = req.Name
		existingProdi.Description = &req.Description
		existingProdi.Parent = fakultas.ID
		if err := repoProdi.UpdateFakultas(existingProdi); err != nil {
			return nil, err
		}
		Prodi = *existingProdi
	} else {
		// Jika tidak ada, buat baru
		Prodi.Name = req.Name
		Prodi.IDNumber = &req.IDNumber
		Prodi.Description = &req.Description
		Prodi.Parent = fakultas.ID

		if err := repoProdi.AddNewFakultas(&Prodi); err != nil {
			return nil, err
		}

		context := model.MdlContext{
			ContextLevel: 40, // Level for course category
			InstanceID:   Prodi.ID,
		}

		if err := repoContext.Create(nil, &context); err != nil {
			return nil, err
		}
	}

	return &Prodi, nil
}

// GetProdi retrieves all Prodi from the database
func (s *MoodleProdiService) GetProdi(id string, db *gorm.DB) ([]model.MdlCourseCategory, error) {
	var repoProdi = repository.NewMoodleFakultasRepository(db)

	var Prodi []model.MdlCourseCategory

	Prodi, err := repoProdi.GetAllProdi(id)
	if err != nil {
		return nil, err
	}

	return Prodi, nil
}

// BatchProdiSync sync all Prodi from all perguruan tinggi and returns a list of errors if any
func (s *MoodleProdiService) BatchProdiSync(req []response.MoodleProdiRequest, db *gorm.DB) []error {
	errs := []error{}
	for _, config := range req {
		_, err := s.AddProdi(config, db)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return []error{}
}
