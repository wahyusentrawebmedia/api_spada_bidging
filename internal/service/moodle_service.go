package service

import (
	"errors"
)

type MoodleService struct{}

func NewMoodleService() *MoodleService {
	return &MoodleService{}
}

// UpdateUserPassword mengupdate password user moodle (dummy, implementasi asli sesuai kebutuhan)
func (s *MoodleService) UpdateUserPassword(userID int, password string) error {
	// TODO: Implementasi update password ke database Moodle
	if userID == 0 || password == "" {
		return errors.New("user_id dan password wajib diisi")
	}
	// Contoh: update ke DB, atau panggil API Moodle
	return nil
}
