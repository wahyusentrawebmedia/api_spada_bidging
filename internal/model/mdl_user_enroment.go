package model

type UserEnrolment struct {
	ID           int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Status       int64 `gorm:"not null;default:0" json:"status"`
	EnrolID      int64 `gorm:"not null;column:enrolid" json:"enrolid"`
	UserID       int64 `gorm:"not null;column:userid" json:"userid"`
	TimeStart    int64 `gorm:"not null;default:0;column:timestart" json:"timestart"`
	TimeEnd      int64 `gorm:"not null;default:2147483647;column:timeend" json:"timeend"`
	ModifierID   int64 `gorm:"not null;default:0;column:modifierid" json:"modifierid"`
	TimeCreated  int64 `gorm:"not null;default:0;column:timecreated" json:"timecreated"`
	TimeModified int64 `gorm:"not null;default:0;column:timemodified" json:"timemodified"`
}

// TableName overrides the table name used by UserEnrolment to `mdl_user_enrolments`
func (UserEnrolment) TableName() string {
	return "mdl_user_enrolments"
}
