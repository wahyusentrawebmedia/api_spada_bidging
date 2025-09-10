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

type ParameterUser struct {
	IdNumberGroup string
	TypeUser      string // dosen atau mahasiswa
	Page          int
	Limit         int
}

// FetchAllUsersWithPagination retrieves all users with pagination
func (s *UserService) FetchAllUsersWithPagination(
	db *gorm.DB,
	param ParameterUser,
) ([]model.MdlUser, error) {
	parameter := repository.ParameterUser{}
	repo := repository.NewUserRepository(db)

	if param.TypeUser != "" {
		if param.TypeUser != "dosen" && param.TypeUser != "mahasiswa" {
			return nil, errors.New("Tipe user tidak valid: " + param.TypeUser)
		}
		parameter.TypeUser = param.TypeUser
	}

	if param.IdNumberGroup != "" {
		repoGroups := repository.NewGroupsRepository(db)
		groups, err := repoGroups.GetByIDNumber(param.IdNumberGroup)
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
func (s *UserService) SyncUserBatchDosenMahasiswaMakul(c *utils.CustomContext, db *gorm.DB, users model.DosenMahasiwaSyncRequest, kodeMakul string) ([]model.UserSyncResponse, error) {
	var results []model.UserSyncResponse

	for _, user := range users.Mahasiswa {
		resp, err := s.SyncUserDosen(c, db, &user, "mahasiswa")
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
		resp, err := s.SyncUserDosen(c, db, &user, "dosen")
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
func (s *UserService) SyncUserBatchDosenMahasiswa(c *utils.CustomContext, db *gorm.DB, users model.DosenMahasiwaSyncRequest) ([]model.UserSyncResponse, error) {
	var results []model.UserSyncResponse

	for _, user := range users.Mahasiswa {
		resp, err := s.SyncUser(c, repository.NewUserRepository(db), &user)
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
		resp, err := s.SyncUserDosen(c, db, &user, "dosen")
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

func (s *UserService) Role(c *utils.CustomContext, db *gorm.DB, userType string) (*model.Role, error) {
	repoRole := repository.NewMoodleRoleRepository(db)

	var role *model.Role
	var err error

	switch userType {
	case "dosen":
		role, err = repoRole.GetByID(c.Context(), 3) // Assuming 3 is the ID for editingteacher
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Tidak bisa mendapatkan role " + userType)
		}

		if role == nil {
			newRole := model.Role{
				Name:        "Editing Teacher",
				ShortName:   "editingteacher",
				Description: "A teacher with editing rights",
				SortOrder:   0,
				Archetype:   "editingteacher",
			}
			err := repoRole.Create(c.Context(), &newRole)
			if err != nil {
				return nil, errors.New("Tidak bisa membuat role " + userType)
			}
			role = &newRole
		}
	case "mahasiswa":
		role, err = repoRole.GetByID(c.Context(), 5) // Assuming 5 is the ID for student
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Tidak bisa mendapatkan role " + userType)
		}
		if role == nil {
			newRole := model.Role{
				Name:        "Student",
				ShortName:   "student",
				Description: "A student",
				SortOrder:   0,
				Archetype:   "student",
			}
			err := repoRole.Create(c.Context(), &newRole)
			if err != nil {
				return nil, errors.New("Tidak bisa membuat role " + userType)
			}
			role = &newRole
		}
	}
	return role, nil
}

func (s *UserService) ContextSystem(c *utils.CustomContext, db *gorm.DB, userType string) (*model.MdlContext, error) {
	repoContext := repository.NewMoodleContextRepository(db)

	// check context ada tidak untuk system atau course berdasarkan userType
	var contextLevel int
	var instanceID int64

	switch userType {
	case "dosen":
		contextLevel = 10 // system context
		instanceID = 0
	case "mahasiswa":
		contextLevel = 10 // system context, bisa diubah jika ada kebutuhan khusus
		instanceID = 0
	default:
		return nil, errors.New("User type tidak dikenali: " + userType)
	}

	contextSystem, err := repoContext.GetByInstanceIDAndLevel(c.Context(), contextLevel)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Tidak bisa mendapatkan context untuk user " + userType)
	}

	// create context system if not exists
	if contextSystem == nil {
		newContext := model.MdlContext{
			ContextLevel: int64(contextLevel),
			InstanceID:   instanceID,
			// Tambahkan field lain jika diperlukan
		}
		err := repoContext.Create(c.Context(), &newContext)
		if err != nil {
			return nil, errors.New("Tidak bisa membuat context untuk user " + userType)
		}
		contextSystem = &newContext
	}
	return contextSystem, nil
}

func (s *UserService) RegisterToCohort(c *utils.CustomContext, db *gorm.DB, userID int64, userType string) error {
	repoCohort := repository.NewMoodleCohortRepository(db)
	repoCohortMember := repository.NewMoodleCohortMemberRepository(db)

	contextSystem, err := s.ContextSystem(c, db, userType)

	if err != nil {
		return err
	}

	// check cohort ada tidak untuk dosen/mahasiswa kalau tidak ada buatkan
	var cohortIDNumber string
	var cohortName string
	var cohortDesc string

	switch userType {
	case "dosen":
		cohortIDNumber = "dosen"
		cohortName = "Dosen"
		cohortDesc = "Cohort for Dosen"
	case "mahasiswa":
		cohortIDNumber = "mahasiswa"
		cohortName = "Mahasiswa"
		cohortDesc = "Cohort for Mahasiswa"
	default:
		return errors.New("User type tidak dikenali: " + userType)
	}

	cohort, err := repoCohort.GetCohortByIDNumber(cohortIDNumber)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("Tidak bisa mendapatkan cohort untuk user " + userType)
	}

	if cohort == nil {
		newCohort := model.Cohort{
			Name:         cohortName,
			IDNumber:     utils.StringPtr(cohortIDNumber),
			Description:  utils.StringPtr(cohortDesc),
			ContextID:    contextSystem.ID,
			TimeCreated:  time.Now().Unix(),
			TimeModified: time.Now().Unix(),
		}
		err := repoCohort.Create(c.Context(), &newCohort)
		if err != nil {
			return errors.New("Tidak bisa membuat cohort untuk user " + userType)
		}
		cohort = &newCohort
	}

	// check apakah user sudah ada di cohort member
	cohortMember, err := repoCohortMember.GetByCohortIDAndUserId(c.Context(), cohort.ID, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("Tidak bisa mendapatkan cohort member untuk user " + userType)
	}

	if cohortMember == nil {
		newCohortMember := model.CohortMember{
			CohortID:  cohort.ID,
			UserID:    userID,
			TimeAdded: time.Now().Unix(),
		}
		err := repoCohortMember.Create(c.Context(), &newCohortMember)
		if err != nil {
			return errors.New("Tidak bisa menambahkan user ke cohort " + cohortName + " " + userType)
		}
	}

	return nil
}

func (s *UserService) AddRoleAssignment(c *utils.CustomContext, db *gorm.DB, userID int64, userType string) error {
	repoRoleAssignment := repository.NewRoleAssignmentRepository(db)

	// check role ada tidak editingteacher kalau tidak ada buatkan
	role, err := s.Role(c, db, userType)

	if err != nil {
		return err
	}

	contextSystem, err := s.ContextSystem(c, db, userType)

	if err != nil {
		return err
	}

	newRoleAssignment := model.RoleAssignment{
		RoleID:       role.ID,
		ContextID:    contextSystem.ID,
		UserID:       userID,
		TimeModified: time.Now().Unix(),
		ModifierID:   2, // Assuming 2 is the ID of the admin user
		Component:    "moodle",
		ItemID:       0,
		SortOrder:    0,
	}
	err = repoRoleAssignment.Create(c.Context(), &newRoleAssignment)
	if err != nil {
		return errors.New("Tidak bisa menambahkan role assignment untuk user " + userType)
	}

	return nil
}

func (s *UserService) SyncUserDosen(c *utils.CustomContext, db *gorm.DB, user *model.UserSyncRequest, userType string) (*model.UserSyncResponse, error) {
	if userType != "dosen" && userType != "mahasiswa" {
		return nil, errors.New("Tipe user tidak valid: " + userType)
	}

	repoMahasiswa := repository.NewUserRepository(db)
	repoApiMoodle := repository.NewApiModel(c.GetEndpoint())
	repoCohort := repository.NewMoodleCohortRepository(db)

	// check context ada tidak untuk system kalau tidak ada buatkan
	contextSystem, err := s.ContextSystem(c, db, userType)

	if err != nil {
		return nil, err
	}

	// check cohort ada tidak untuk dosen kalau tidak ada buatkan
	if userType == "dosen" {
		cohort, err := repoCohort.GetCohortByIDNumber("dosen")
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Tidak bisa mendapatkan cohort untuk user " + userType)
		}

		if cohort == nil {
			newCohort := model.Cohort{
				Name:         "Dosen",
				IDNumber:     utils.StringPtr("dosen"),
				Description:  utils.StringPtr("Cohort for Dosen"),
				ContextID:    contextSystem.ID,
				TimeCreated:  time.Now().Unix(),
				TimeModified: time.Now().Unix(),
			}
			err := repoCohort.Create(c.Context(), &newCohort)
			if err != nil {
				return nil, errors.New("Tidak bisa membuat cohort untuk user " + user.Username)
			}
			cohort = &newCohort
		}
	}

	// check username ada tidak di user existing
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

		s.RegisterToCohort(c, db, int64(userExists.ID), userType)
		s.AddRoleAssignment(c, db, int64(userExists.ID), userType)

		if err != nil {
			return &model.UserSyncResponse{
				Action:   true,
				Username: userExists.Username,
				Password: userExists.Password,
				Pesan:    "Sinkronisasi Gagal : " + err.Error(),
				IdSpada:  userExists.ID,
			}, nil
		} else {
			return &model.UserSyncResponse{
				Action:   false,
				Username: userExists.Username,
				Password: userExists.Password,
				Pesan:    "Sinkronisasi Berhasil",
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

		s.RegisterToCohort(c, db, int64(userExists.ID), userType)
		s.AddRoleAssignment(c, db, int64(userExists.ID), userType)

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

	// INSERT INTO mdl_role_assignments (roleid, contextid, userid, timemodified, modifierid)
	// VALUES (@role_id, @context_id, @user_id, UNIX_TIMESTAMP(), 2)
	// ON DUPLICATE KEY UPDATE timemodified = UNIX_TIMESTAMP();
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
