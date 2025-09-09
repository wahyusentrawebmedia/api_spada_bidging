package model

type Cohort struct {
	ID                int64   `gorm:"column:id;primaryKey;autoIncrement"`
	ContextID         int64   `gorm:"column:contextid;not null;index"`
	Name              string  `gorm:"column:name;size:254;not null;default:''"`
	IDNumber          *string `gorm:"column:idnumber;size:100"`
	Description       *string `gorm:"column:description;type:longtext"`
	DescriptionFormat int8    `gorm:"column:descriptionformat;not null"`
	Visible           int8    `gorm:"column:visible;not null;default:1"`
	Component         string  `gorm:"column:component;size:100;not null;default:''"`
	TimeCreated       int64   `gorm:"column:timecreated;not null"`
	TimeModified      int64   `gorm:"column:timemodified;not null"`
	Theme             *string `gorm:"column:theme;size:50"`
}

// TableName overrides the table name used by Cohort to `mdl_cohort`
func (Cohort) TableName() string {
	return "mdl_cohort"
}
