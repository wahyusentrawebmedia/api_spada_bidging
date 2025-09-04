package service

import (
	"errors"
	"time"

	"api/spada/internal/model"
	"api/spada/internal/repository"
	"api/spada/internal/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

// FetchAllUsersWithPagination retrieves all users with pagination
func (s *UserService) FetchAllUsersWithPagination(repo *repository.UserRepository, page, limit int) ([]model.MdlUser, error) {
	users, err := repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Example CRUD method
func (s *UserService) GetUserByID(id int) (interface{}, error) {
	// Implement logic to get user by ID
	return nil, nil
}

// SyncUser synchronizes user data from an external source
func (s *UserService) SyncUser(c *utils.CustomContext, repo *repository.UserRepository, user *model.UserSyncRequest) (*model.UserSyncResponse, error) {
	repoMahasiswa := repo
	repoApiMoodle := repository.NewApiModel(c.GetEndpoint())

	userExists, err := repoMahasiswa.GetUserByUsername(user.Username)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Tidak bisa mendapatkan user untuk user " + user.Username)
	}

	if userExists != nil {

		userExists.Password = user.Password
		userExists.FirstName = user.FirstName
		userExists.LastName = user.LastName
		userExists.Email = user.Email
		userExists.TimeModified = time.Now().Unix()

		err := repoMahasiswa.UpdateUser(userExists)

		if err != nil {
			return &model.UserSyncResponse{
				Action:   false,
				Username: userExists.Username,
				Password: userExists.Password,
				Pesan:    "Sinkronisasi Berhasil",
			}, nil
		} else {
			return &model.UserSyncResponse{
				Action:   true,
				Username: userExists.Username,
				Password: userExists.Password,
				Pesan:    "Sinkronisasi Gagal",
				IdSpada:  userExists.ID,
			}, nil
		}
	} else {
		simpan := model.MdlUser{
			Auth:         "manual",
			Confirmed:    1,
			MnethostID:   1,
			Username:     user.Username,
			Password:     repoApiMoodle.HashingPassword(user.Password), // You may want to hash the password here
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			Email:        user.Email,
			City:         "Medan",
			Country:      "ID",
			Lang:         "id",
			Timezone:     "Asia/Jakarta",
			TimeCreated:  time.Now().Unix(),
			TimeModified: time.Now().Unix(),
		}

		err := repoMahasiswa.CreateUser(&simpan)

		if err != nil {
			return &model.UserSyncResponse{
				Action:   false,
				Username: user.Username,
				Password: user.Password,
				Pesan:    "Gagal menyimpan user " + err.Error(),
			}, nil
		} else {
			return &model.UserSyncResponse{
				Action:   true,
				Username: user.Username,
				Password: user.Password,
				Pesan:    "Berhasil menyimpan user",
				IdSpada:  simpan.ID,
			}, nil
		}
	}
}

// ResetPassword resets the password for a user
func (s *UserService) ResetPassword(repo *repository.UserRepository, ids []int) (interface{}, error) {
	repoUsers := repo
	var updatedUsers []model.MdlUser
	var newPasswords = make(map[int]string)

	for _, id := range ids {
		user, err := repoUsers.GetUserBgyID(id)
		if err != nil || user == nil {
			continue
		}

		newPassword := generateRandomString(10)
		hashedPassword, err := HashPassword(newPassword)
		if err != nil {
			continue
		}

		user.Password = hashedPassword
		err = repoUsers.UpdateUser(user)
		if err == nil {
			updatedUsers = append(updatedUsers, *user)
			newPasswords[id] = newPassword
		}
	}

	return map[string]interface{}{
		"updated_users": updatedUsers,
		"new_passwords": newPasswords,
	}, nil
}

// generateRandomString generates a random string of given length
func generateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[time.Now().UnixNano()%int64(len(letters))]
	}
	return string(b)
}

// HashPassword hashes the password using bcrypt
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}
