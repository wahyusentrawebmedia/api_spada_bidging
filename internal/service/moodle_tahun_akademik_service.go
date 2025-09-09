package service

import (
	"api/spada/internal/model"
	"api/spada/internal/repository"
	"api/spada/internal/response"
	"api/spada/internal/utils"

	"gorm.io/gorm"
)

type MoodleTahunAkademikService struct {
}

func NewMoodleTahunAkademikService() *MoodleTahunAkademikService {
	return &MoodleTahunAkademikService{}
}

// AddTahunAkademik adds a new TahunAkademik to the database
func (s *MoodleTahunAkademikService) AddTahunAkademik(req response.MoodleTahunAkademikRequest, db *gorm.DB) (*model.MdlCourseCategory, error) {
	var repoTahunAkademik = repository.NewMoodleFakultasRepository(db)
	var repoContext = repository.NewMoodleContextRepository(db)
	var repoCohort = repository.NewMoodleCohortRepository(db)
	var repoUserInfoField = repository.NewMoodleUserInfoFieldRepository(db)
	var repoUserInfoData = repository.NewMoodleUserInfoDataRepository(db)

	// Cek apakah TahunAkademik dengan IDNumber yang sama sudah ada
	existingTahunAkademik, err := repoTahunAkademik.GetFakultasByIDNumber(req.IDNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	prodi, err := repoTahunAkademik.GetFakultasByIDNumber(req.Parent)
	if err != nil {
		return nil, err
	}

	var TahunAkademik model.MdlCourseCategory

	if existingTahunAkademik != nil && existingTahunAkademik.ID > 0 {
		// Jika ada, update data TahunAkademik
		existingTahunAkademik.Name = req.Name
		existingTahunAkademik.Description = &req.Description
		existingTahunAkademik.Parent = prodi.ID

		if err := repoTahunAkademik.UpdateFakultas(existingTahunAkademik); err != nil {
			return nil, err
		}
		TahunAkademik = *existingTahunAkademik
	} else {
		// Jika tidak ada, buat baru
		TahunAkademik.Name = req.Name
		TahunAkademik.IDNumber = &req.IDNumber
		TahunAkademik.Description = &req.Description
		TahunAkademik.Parent = prodi.ID

		if err := repoTahunAkademik.AddNewFakultas(&TahunAkademik); err != nil {
			return nil, err
		}

		context := model.MdlContext{
			ContextLevel: 40, // Level for course category
			InstanceID:   TahunAkademik.ID,
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

	// Tambah field custom jika belum ada
	field, err := repoUserInfoField.GetByShortName(nil, "tahun_akademik")
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if field == nil || field.ID == 0 {
		newField := model.UserInfoField{
			ShortName:   "tahun_akademik",
			Name:        "Tahun Akademik",
			DataType:    "text",
			Description: utils.PtrString("Field untuk menyimpan tahun akademik pengguna"),
		}
		if err := repoUserInfoField.Create(nil, &newField); err != nil {
			return nil, err
		}
		field = &newField
	}

	// Update data user: set tahun_akademik for all users if not already set
	users, err := repoUserInfoData.GetAllUsersWithoutField(nil, field.ID)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		userInfoData := model.UserInfoData{
			UserID:  user.ID,
			FieldID: field.ID,
			Data:    req.Name, // or req.IDNumber if you want to store the IDNumber
		}
		if err := repoUserInfoData.Create(nil, &userInfoData); err != nil {
			return nil, err
		}
	}

	return &TahunAkademik, nil
}

// GetTahunAkademik retrieves all TahunAkademik from the database
func (s *MoodleTahunAkademikService) GetTahunAkademik(db *gorm.DB) ([]model.MdlCourseCategory, error) {
	var repoTahunAkademik = repository.NewMoodleFakultasRepository(db)

	var TahunAkademik []model.MdlCourseCategory

	TahunAkademik, err := repoTahunAkademik.GetAllFakultas()
	if err != nil {
		return nil, err
	}

	return TahunAkademik, nil
}

// BatchTahunAkademikSync sync all TahunAkademik from all perguruan tinggi and returns a list of errors if any
func (s *MoodleTahunAkademikService) BatchTahunAkademikSync(req []response.MoodleTahunAkademikRequest, db *gorm.DB) []error {
	var errs []error
	for _, config := range req {
		_, err := s.AddTahunAkademik(config, db)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}
