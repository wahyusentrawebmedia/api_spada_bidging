package model

type MdlScale struct {
	ID                int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	CourseID          int64  `json:"courseid" gorm:"column:courseid"`
	UserID            int64  `json:"userid" gorm:"column:userid"`
	Name              string `json:"name" gorm:"column:name"`
	Scale             string `json:"scale" gorm:"column:scale"`
	Description       string `json:"description" gorm:"column:description"`
	DescriptionFormat int16  `json:"descriptionformat" gorm:"column:descriptionformat"`
	TimeModified      int64  `json:"timemodified" gorm:"column:timemodified"`
}

func (MdlScale) TableName() string {
	return "mdl_scale"
}
