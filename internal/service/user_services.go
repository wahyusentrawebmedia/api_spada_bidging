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
func (s *UserService) FetchAllUsersWithPagination(
	idNumberGroup string,
	db *gorm.DB,
	page,
	limit int,
) ([]model.MdlUser, error) {
	parameter := repository.ParameterUser{}
	repo := repository.NewUserRepository(db)

	if idNumberGroup != "" {
		repoGroups := repository.NewGroupsRepository(db)
		groups, err := repoGroups.GetByIDNumber(idNumberGroup)
		if err != nil {
			return nil, err
		}
		parameter.IDGrup = int(groups.ID)
	}

	users, err := repo.GetAllUsers(parameter)
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

// GetUserByUsername retrieves a user by their username
func (s *UserService) GetUserByUsername(repo *repository.UserRepository, username string) (*model.MdlUser, error) {
	user, err := repo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return &model.MdlUser{}, nil
	}
	return user, nil
}

// SyncUserBatchDosenMahasiswaMakul synchronizes a batch of dosen and mahasiswa users for a specific makul
func (s *UserService) SyncUserBatchDosenMahasiswaMakul(c *utils.CustomContext, repo *repository.UserRepository, users model.DosenMahasiwaSyncRequest, kodeMakul string) ([]model.UserSyncResponse, error) {
	var results []model.UserSyncResponse

	for _, user := range users.Mahasiswa {
		resp, err := s.SyncUser(c, repo, &user)
		if err != nil {
			results = append(results, model.UserSyncResponse{
				Action:   false,
				Username: user.Username,
				Password: user.Password,
				Pesan:    "Sinkronisasi Gagal: " + err.Error(),
			})
		} else {
			results = append(results, *resp)
		}
	}

	for _, user := range users.Dosen {
		resp, err := s.SyncUser(c, repo, &user)
		if err != nil {
			results = append(results, model.UserSyncResponse{
				Action:   false,
				Username: user.Username,
				Password: user.Password,
				Pesan:    "Sinkronisasi Gagal: " + err.Error(),
			})
		} else {
			results = append(results, *resp)
		}
	}
	return results, nil
}

// SyncUserBatchDosenMahasiswa synchronizes a batch of dosen and mahasiswa users
func (s *UserService) SyncUserBatchDosenMahasiswa(c *utils.CustomContext, repo *repository.UserRepository, users model.DosenMahasiwaSyncRequest) ([]model.UserSyncResponse, error) {
	var results []model.UserSyncResponse

	for _, user := range users.Mahasiswa {
		resp, err := s.SyncUser(c, repo, &user)
		if err != nil {
			results = append(results, model.UserSyncResponse{
				Action:   false,
				Username: user.Username,
				Password: user.Password,
				Pesan:    "Sinkronisasi Gagal: " + err.Error(),
			})
		} else {
			results = append(results, *resp)
		}
	}

	for _, user := range users.Dosen {
		resp, err := s.SyncUser(c, repo, &user)
		if err != nil {
			results = append(results, model.UserSyncResponse{
				Action:   false,
				Username: user.Username,
				Password: user.Password,
				Pesan:    "Sinkronisasi Gagal: " + err.Error(),
			})
		} else {
			results = append(results, *resp)
		}
	}
	return results, nil
}

// SyncUserBatch synchronizes a batch of users
func (s *UserService) SyncUserBatch(c *utils.CustomContext, repo *repository.UserRepository, users []model.UserSyncRequest) ([]model.UserSyncResponse, error) {
	var results []model.UserSyncResponse

	for _, user := range users {
		resp, err := s.SyncUser(c, repo, &user)
		if err != nil {
			results = append(results, model.UserSyncResponse{
				Action:   false,
				Username: user.Username,
				Password: user.Password,
				Pesan:    "Sinkronisasi Gagal: " + err.Error(),
			})
		} else {
			results = append(results, *resp)
		}
	}
	return results, nil
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

// ChangePassword changes the password for a user
func (s *UserService) ChangePassword(c *utils.CustomContext, repo *repository.UserRepository, username, oldPassword, newPassword string) error {
	repoUsers := repo

	repoApiMoodle := repository.NewApiModel(c.GetEndpoint())

	user, err := repoUsers.GetUserByUsername(username)
	if err != nil {
		return err
	}

	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
	// 	return errors.New("Old password is incorrect")
	// }

	hashedNewPassword := repoApiMoodle.HashingPassword(newPassword)
	if err != nil {
		return err
	}

	user.Password = hashedNewPassword
	return repoUsers.UpdateUser(user)
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

// // GetUserByGroupMembershipID retrieves a user by their group membership ID
// func (s *UserService) GetUserByGroupMembershipID(groupMemberID int64) (*model.MdlUser, error) {
// 	groupMember, err := repoGroupMember.GetByID(context.Background(), groupMemberID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if groupMember == nil {
// 		return nil, errors.New("Group member not found")
// 	}

// 	user, err := repoUser.GetUserBgyID(int(groupMember.UserID))
// 	if err != nil {
// 		return nil, err
// 	}
// 	if user == nil {
// 		return nil, errors.New("User not found")
// 	}

// 	return user, nil
// }
