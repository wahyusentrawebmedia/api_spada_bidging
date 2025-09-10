package model

type MdlGroups struct {
	ID                int64   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CourseID          int64   `gorm:"column:courseid;not null" json:"courseid"`
	IDNumber          string  `gorm:"column:idnumber;size:100;not null;default:''" json:"idnumber"`
	Name              string  `gorm:"column:name;size:254;not null;default:''" json:"name"`
	Description       *string `gorm:"column:description;type:longtext" json:"description,omitempty"`
	DescriptionFormat int8    `gorm:"column:descriptionformat;not null;default:0" json:"descriptionformat"`
	EnrolmentKey      *string `gorm:"column:enrolmentkey;size:50" json:"enrolmentkey,omitempty"`
	Picture           int64   `gorm:"column:picture;not null;default:0" json:"picture"`
	Visibility        int8    `gorm:"column:visibility;not null;default:0" json:"visibility"`
	Participation     int8    `gorm:"column:participation;not null;default:1" json:"participation"`
	TimeCreated       int64   `gorm:"column:timecreated;not null;default:0" json:"timecreated"`
	TimeModified      int64   `gorm:"column:timemodified;not null;default:0" json:"timemodified"`
}

// TableName overrides the table name used by GORM.
func (MdlGroups) TableName() string {
	return "mdl_groups"
}
