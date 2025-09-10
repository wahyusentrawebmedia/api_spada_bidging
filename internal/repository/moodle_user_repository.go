package repository

import (
	"api/spada/internal/model"

	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

type UserRepository struct {
	db *gorm.DB
}

type ParameterUser struct {
	IDGrup   int
	TypeUser string
	IdMakul  string
}

func (r *UserRepository) GetAllUsers(parameter ParameterUser) ([]model.MdlUser, error) {
	var users []model.MdlUser

	query := r.db.Debug().Model(&model.MdlUser{})

	if parameter.IDGrup != 0 {
		query = query.
			Joins("JOIN mdl_groups_members gm ON gm.userid = mdl_user.id").
			Where("gm.groupid = ?", parameter.IDGrup)
	}

	if parameter.TypeUser != "" {
		switch parameter.TypeUser {
		case "dosen":
			subQuery := r.db.Model(&model.MdlUser{}).
				Joins("JOIN mdl_role_assignments ra ON mdl_user.id = ra.userid").
				Joins("JOIN mdl_role r ON ra.roleid = r.id").
				Where("r.shortname = ?", "editingteacher").
				Select("mdl_user.id")

			query = query.Where("mdl_user.id IN (?)", subQuery)
		case "mahasiswa":
			subQuery := r.db.Model(&model.MdlUser{}).
				Joins("JOIN mdl_role_assignments ra ON mdl_user.id = ra.userid").
				Joins("JOIN mdl_role r ON ra.roleid = r.id").
				Where("r.shortname = ?", "student").
				Select("mdl_user.id")

			query = query.Where("mdl_user.id IN (?)", subQuery)
		}
	}

	if parameter.IdMakul != "" {
		subQuery := r.db.Model(&model.MdlUser{}).
			Joins("JOIN mdl_user_enrolments ue ON ue.userid = mdl_user.id").
			Joins("JOIN mdl_enrol e ON ue.enrolid = e.id").
			Joins("JOIN mdl_course c ON e.courseid = c.id").
			Where("c.idnumber = ?", parameter.IdMakul).
			Select("mdl_user.id")

		query = query.Where("mdl_user.id IN (?)", subQuery)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID retrieves a user by their ID
func (r *UserRepository) GetUserBgyID(id int) (*model.MdlUser, error) {
	var user model.MdlUser
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByUsername retrieves a user by their username
func (r *UserRepository) GetUserByUsername(username string) (*model.MdlUser, error) {
	var user model.MdlUser
	err := r.db.Debug().Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates a user's information
func (r *UserRepository) UpdateUser(user *model.MdlUser) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

// CreateUser creates a new user
func (r *UserRepository) CreateUser(user *model.MdlUser) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
