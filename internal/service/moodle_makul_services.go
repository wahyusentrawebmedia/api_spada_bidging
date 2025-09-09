package service

import (
	"api/spada/internal/model"
	"api/spada/internal/repository"
	"api/spada/internal/response"
	"api/spada/internal/utils"
	"time"

	"gorm.io/gorm"
)

type MoodleMakulService struct {
	// DB *gorm.DB
}

func NewMoodleMakulService() *MoodleMakulService {
	return &MoodleMakulService{}
}

// SyncMakulAll syncs multiple courses (makul) in Moodle based on the provided request data.
func (s *MoodleMakulService) SyncMakulAll(reqs []response.MoodleMakulRequest, parent string, db *gorm.DB) []error {
	var errs []error
	for _, req := range reqs {
		if err := s.SyncMakul(req, parent, db); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// SyncMakul syncs a course (makul) in Moodle based on the provided request data.
func (s *MoodleMakulService) SyncMakul(req response.MoodleMakulRequest, parent string, db *gorm.DB) error {
	repoCategories := repository.NewMoodleFakultasRepository(db)
	repoCourse := repository.NewMoodleCourseRepository(db)
	repoGroups := repository.NewGroupsRepository(db)

	idnumber := req.KodeMK + "_" + req.Tahun + "_" + req.Semester + "_" + req.Kelas

	// getParent
	parentCategories, err := repoCategories.GetFakultasByIDNumber(parent)
	if err != nil {
		return err
	}

	// create or update course
	repoCourseData, err := repoCourse.GetByIDNumber(idnumber)
	if err != nil {
		return err
	}

	if repoCourseData != nil && repoCourseData.ID > 0 {
		// Update existing course
		repoCourseData.FullName = req.NamaMK + " - " + req.Kelas + " - " + req.Tahun
		repoCourseData.ShortName = req.KodeMK + " - " + req.Kelas + " - " + req.Tahun
		repoCourseData.Category = parentCategories.ID
		// repoCourseData.Summary = "Kelas " + req.Kelas + " - " + req.NamaMK
		repoCourseData.Format = "topics"
		// repoCourseData.StartDate = 1693526400 // strtotime('2023-09-01')
		// repoCourseData.EndDate = 1703980800   // strtotime('2023-12-31')
		if err := repoCourse.Update(nil, repoCourseData); err != nil {
			return err
		}
	} else {
		// Create new course
		newCourse := &model.Course{
			FullName:  req.NamaMK + " - " + req.Kelas + " - " + req.Tahun,
			ShortName: req.KodeMK + " - " + req.Kelas + " - " + req.Tahun,
			// Summary:   "Kelas " + req.Kelas + " - " + req.NamaMK,
			Format: "topics",
			// StartDate: strtotime('2023-09-01'),
			// EndDate:   strtotime('2023-12-31'),
			IDNumber: idnumber,
			Category: parentCategories.ID,
			Visible:  1,
		}
		if err := repoCourse.Create(nil, newCourse); err != nil {
			return err
		}

		repoCourseData = newCourse
	}

	// create or update groups
	groupName := req.KodeMK + " - " + req.Kelas + " - " + req.Tahun
	idNumberGroup := idnumber + "_GRP"
	group, err := repoGroups.GetByIDNumber(idNumberGroup)
	if err != nil {
		return err
	}

	if group != nil && group.ID > 0 {
		// Update existing group
		group.Name = groupName
		if err := repoGroups.Update(nil, group); err != nil {
			return err
		}
	} else {
		// Create new group
		newGroup := &model.MdlGroups{
			CourseID:          repoCourseData.ID,
			Name:              groupName,
			IDNumber:          idNumberGroup,
			Description:       utils.PtrString("Group for " + req.NamaMK + " - " + req.Kelas + " - " + req.Tahun),
			DescriptionFormat: 1,
			TimeCreated:       time.Now().Unix(),
			TimeModified:      time.Now().Unix(),
		}
		if err := repoGroups.Create(nil, newGroup); err != nil {
			return err
		}
	}

	return nil
}
