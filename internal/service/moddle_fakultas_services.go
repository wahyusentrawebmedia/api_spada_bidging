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

	// Cek apakah fakultas dengan IDNumber yang sama sudah ada
	existingFakultas, err := repoFakultas.GetFakultasByIDNumber(req.IDNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	var fakultas model.MdlCourseCategory

	if existingFakultas != nil && existingFakultas.ID > 0 {
		// Jika ada, update data fakultas
		existingFakultas.Name = req.Name
		existingFakultas.Description = &req.Description
		if err := repoFakultas.UpdateFakultas(existingFakultas); err != nil {
			return nil, err
		}
		fakultas = *existingFakultas
	} else {
		// Jika tidak ada, buat baru
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
	}

	return &fakultas, nil
}

// GetFakultas retrieves all fakultas from the database
func (s *MoodleFakultasService) GetFakultas(db *gorm.DB) ([]model.MdlCourseCategory, error) {
	var repoFakultas = repository.NewMoodleFakultasRepository(db)

	var fakultas []model.MdlCourseCategory

	fakultas, err := repoFakultas.GetAllFakultas()
	if err != nil {
		return nil, err
	}

	return fakultas, nil
}

// BatchFakultasSync sync all fakultas from all perguruan tinggi and returns a list of errors if any
func (s *MoodleFakultasService) BatchFakultasSync(req []response.MoodleFakultasRequest, db *gorm.DB) []error {
	var errs []error
	for _, config := range req {
		_, err := s.AddFakultas(config, db)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}
