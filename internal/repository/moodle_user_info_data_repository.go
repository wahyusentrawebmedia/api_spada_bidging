package repository

import (
	"api/spada/internal/model"
	"context"

	"gorm.io/gorm"
)

type MoodleUserInfoDataRepository interface {
	Create(ctx context.Context, data *model.UserInfoData) error
	GetByID(ctx context.Context, id int64) (*model.UserInfoData, error)
	Update(ctx context.Context, data *model.UserInfoData) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*model.UserInfoData, error)
	GetAllUsersWithoutField(ctx context.Context, fieldID int64) ([]model.MdlUser, error)
}

type moodleUserInfoDataRepository struct {
	db *gorm.DB
}

func NewMoodleUserInfoDataRepository(db *gorm.DB) MoodleUserInfoDataRepository {
	return &moodleUserInfoDataRepository{db: db}
}

func (r *moodleUserInfoDataRepository) Create(ctx context.Context, data *model.UserInfoData) error {
	return r.db.WithContext(ctx).Create(data).Error
}

func (r *moodleUserInfoDataRepository) GetByID(ctx context.Context, id int64) (*model.UserInfoData, error) {
	var d model.UserInfoData
	if err := r.db.WithContext(ctx).First(&d, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &d, nil
}

func (r *moodleUserInfoDataRepository) Update(ctx context.Context, data *model.UserInfoData) error {
	return r.db.WithContext(ctx).Save(data).Error
}

func (r *moodleUserInfoDataRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.UserInfoData{}, id).Error
}

func (r *moodleUserInfoDataRepository) List(ctx context.Context) ([]*model.UserInfoData, error) {
	var result []*model.UserInfoData
	if err := r.db.WithContext(ctx).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllUsersWithoutField retrieves all user IDs that do not have an entry for the specified field ID
func (r *moodleUserInfoDataRepository) GetAllUsersWithoutField(ctx context.Context, fieldID int64) ([]model.MdlUser, error) {
	var users []model.MdlUser
	subQuery := r.db.Model(&model.UserInfoData{}).Select("userid").Where("fieldid = ?", fieldID)
	if err := r.db.WithContext(ctx).Model(&model.MdlUser{}).
		Where("id NOT IN (?)", subQuery).
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
