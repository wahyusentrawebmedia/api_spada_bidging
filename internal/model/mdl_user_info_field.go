package model

type UserInfoField struct {
	ID                int64   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ShortName         string  `gorm:"column:shortname;size:255;not null;default:shortname" json:"shortname"`
	Name              string  `gorm:"column:name;type:longtext;not null" json:"name"`
	DataType          string  `gorm:"column:datatype;size:255;not null;default:''" json:"datatype"`
	Description       *string `gorm:"column:description;type:longtext" json:"description,omitempty"`
	DescriptionFormat int8    `gorm:"column:descriptionformat;not null;default:0" json:"descriptionformat"`
	CategoryID        int64   `gorm:"column:categoryid;not null;default:0" json:"categoryid"`
	SortOrder         int64   `gorm:"column:sortorder;not null;default:0" json:"sortorder"`
	Required          int8    `gorm:"column:required;not null;default:0" json:"required"`
	Locked            int8    `gorm:"column:locked;not null;default:0" json:"locked"`
	Visible           int16   `gorm:"column:visible;not null;default:0" json:"visible"`
	ForceUnique       int8    `gorm:"column:forceunique;not null;default:0" json:"forceunique"`
	Signup            int8    `gorm:"column:signup;not null;default:0" json:"signup"`
	DefaultData       *string `gorm:"column:defaultdata;type:longtext" json:"defaultdata,omitempty"`
	DefaultDataFormat int8    `gorm:"column:defaultdataformat;not null;default:0" json:"defaultdataformat"`
	Param1            *string `gorm:"column:param1;type:longtext" json:"param1,omitempty"`
	Param2            *string `gorm:"column:param2;type:longtext" json:"param2,omitempty"`
	Param3            *string `gorm:"column:param3;type:longtext" json:"param3,omitempty"`
	Param4            *string `gorm:"column:param4;type:longtext" json:"param4,omitempty"`
	Param5            *string `gorm:"column:param5;type:longtext" json:"param5,omitempty"`
}

// TableName overrides the table name used by UserInfoField to `mdl_user_info_field`
func (UserInfoField) TableName() string {
	return "mdl_user_info_field"
}
