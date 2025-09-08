package model

type UserInfoData struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID     int64  `gorm:"column:userid;not null" json:"userid"`
	FieldID    int64  `gorm:"column:fieldid;not null" json:"fieldid"`
	Data       string `gorm:"column:data;type:longtext;not null" json:"data"`
	DataFormat int8   `gorm:"column:dataformat;not null" json:"dataformat"`
}

// TableName overrides the table name used by GORM.
func (UserInfoData) TableName() string {
	return "mdl_user_info_data"
}
