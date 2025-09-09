package service

import (
	"api/spada/internal/model"
	"api/spada/internal/repository"
	"api/spada/internal/response"

	"gorm.io/gorm"
)

type MoodleSemesterService struct {
}

func NewMoodleSemesterService() *MoodleSemesterService {
	return &MoodleSemesterService{}
}

// AddSemester adds a new Semester to the database
func (s *MoodleSemesterService) AddSemester(req response.MoodleSemesterRequest, db *gorm.DB) (*model.MdlCourseCategory, error) {
	var repoSemester = repository.NewMoodleFakultasRepository(db)
	var repoContext = repository.NewMoodleContextRepository(db)
	var repoCohort = repository.NewMoodleCohortRepository(db)

	// getParent
	fakultas, err := repoSemester.GetFakultasByIDNumber(req.Parent)
	if err != nil {
		return nil, err
	}
	// Cek apakah Semester dengan IDNumber yang sama sudah ada
	existingSemester, err := repoSemester.GetFakultasByIDNumber(req.IDNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	var Semester model.MdlCourseCategory

	if existingSemester != nil && existingSemester.ID > 0 {
		// Jika ada, update data Semester
		existingSemester.IDNumber = &req.IDNumber
		existingSemester.Name = req.Name
		existingSemester.Description = &req.Description
		existingSemester.Parent = fakultas.ID
		if err := repoSemester.UpdateFakultas(existingSemester); err != nil {
			return nil, err
		}
		Semester = *existingSemester
	} else {
		// Jika tidak ada, buat baru
		Semester.IDNumber = &req.IDNumber
		Semester.Name = req.Name
		Semester.Description = &req.Description
		Semester.Parent = fakultas.ID

		if err := repoSemester.AddNewFakultas(&Semester); err != nil {
			return nil, err
		}

		context := model.MdlContext{
			ContextLevel: 40, // Level for course category
			InstanceID:   Semester.ID,
		}

		if err := repoContext.Create(nil, &context); err != nil {
			return nil, err
		}
	}

	cohort, err := repoCohort.GetCohortByIDNumber(req.IDNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if cohort != nil && cohort.ID > 0 {
		// Jika cohort sudah ada, update
		cohort.Name = req.Name
		cohort.Description = &req.Description
		if err := repoCohort.Update(nil, cohort); err != nil {
			return nil, err
		}
	} else {
		// Jika cohort belum ada, buat baru
		newCohort := model.Cohort{
			Name:        req.Name,
			IDNumber:    &req.IDNumber,
			Description: &req.Description,
		}
		if err := repoCohort.Create(nil, &newCohort); err != nil {
			return nil, err
		}
	}

	// Update idnumber for mdl_course based on kode_matkul, tahun, and semester
	// updateQuery := `
	// 	UPDATE mdl_course
	// 	SET idnumber = CONCAT(kode_matkul, '_', tahun, '_', semester)
	// 	WHERE idnumber IS NULL OR idnumber = ''
	// `
	// if err := db.Exec(updateQuery).Error; err != nil {
	// 	return nil, err
	// }

	return &Semester, nil
}

// GetSemester retrieves all Semester from the database
func (s *MoodleSemesterService) GetSemester(db *gorm.DB) ([]model.MdlCourseCategory, error) {
	var repoSemester = repository.NewMoodleFakultasRepository(db)

	var Semester []model.MdlCourseCategory

	Semester, err := repoSemester.GetAllFakultas()
	if err != nil {
		return nil, err
	}

	return Semester, nil
}

// BatchSemesterSync sync all Semester from all perguruan tinggi and returns a list of errors if any
func (s *MoodleSemesterService) BatchSemesterSync(req []response.MoodleSemesterRequest, db *gorm.DB) []error {
	var errs []error
	for _, config := range req {
		_, err := s.AddSemester(config, db)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

// GetDetailSemesterByNumberId retrieves a Semester by its IDNumber
func (s *MoodleSemesterService) GetDetailSemesterByNumberId(idnumber string, db *gorm.DB) (*model.MdlCourseCategory, error) {
	var repoSemester = repository.NewMoodleFakultasRepository(db)

	Semester, err := repoSemester.GetFakultasByIDNumber(idnumber)
	if err != nil {
		return nil, err
	}
	if Semester == nil {
		return nil, gorm.ErrRecordNotFound
	}

	return Semester, nil
}
