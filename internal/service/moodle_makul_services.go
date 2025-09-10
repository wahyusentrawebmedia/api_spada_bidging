package service

import (
	"api/spada/internal/model"
	"api/spada/internal/repository"
	"api/spada/internal/response"
	"api/spada/internal/utils"
	"errors"
	"fmt"
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
	repoCategories := repository.NewMoodleFakultasRepository(db) // sesuai semester
	repoCourse := repository.NewMoodleCourseRepository(db)       // untuk mata kuliah
	repoGroups := repository.NewGroupsRepository(db)             // untuk kelas
	repoContext := repository.NewMoodleContextRepository(db)

	idnumber := req.KodeMK + "_" + req.Tahun + "_" + req.Semester + "_" + req.Kelas

	// getParent
	parentCategories, err := repoCategories.GetFakultasByIDNumber(parent)
	if err != nil {
		return errors.New("parent kategori tidak ditemukan")
	}

	// create or update course
	repoCourseData, err := repoCourse.GetByIDNumber(idnumber)
	if err != nil {
		return errors.New("error mencari course dengan idnumber " + idnumber + ": " + err.Error())
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
			return errors.New("error mengupdate course dengan idnumber " + idnumber + ": " + err.Error())
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
			return errors.New("error membuat course dengan idnumber " + idnumber + ": " + err.Error())
		}

		repoCourseData = newCourse
	}

	// create context if not exists
	context, err := repoContext.GetByInstanceIDAndLevel(nil, int(repoCourseData.ID), 50)
	if err != nil {
		return errors.New("error mencari context course dengan instanceid " + fmt.Sprintf("%d", repoCourseData.ID) + ": " + err.Error())
	}
	if context == nil || context.ID == 0 {
		newContext := &model.MdlContext{
			ContextLevel: 50, // CONTEXT_COURSE
			InstanceID:   repoCourseData.ID,
		}
		if err := repoContext.Create(nil, newContext); err != nil {
			return errors.New("error membuat context course dengan instanceid " + fmt.Sprintf("%d", repoCourseData.ID) + ": " + err.Error())
		}
	}

	// create or update groups
	groupName := req.KodeMK + " - " + req.Kelas + " - " + req.Tahun
	idNumberGroup := idnumber + "_GRP"
	group, err := repoGroups.GetByIDNumber(idNumberGroup)
	if err != nil {
		return errors.New("error mencari group dengan idnumber " + idNumberGroup + ": " + err.Error())
	}

	if group != nil && group.ID > 0 {
		group.Name = groupName
		group.TimeModified = time.Now().Unix()

		if err := repoGroups.Update(nil, group); err != nil {
			return errors.New("error mengupdate group dengan idnumber " + idNumberGroup + ":, id : " + fmt.Sprintf("%d", group.ID) + " " + err.Error())
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
			return errors.New("error membuat group dengan idnumber " + idNumberGroup + ": " + err.Error())
		}
	}

	return nil
}
