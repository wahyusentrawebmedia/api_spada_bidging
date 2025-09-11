package model

type Enrol struct {
	ID              int64    `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Enrol           string   `json:"enrol" gorm:"type:varchar(20);not null;default:'';column:enrol"`
	Status          int64    `json:"status" gorm:"type:bigint;not null;default:0;column:status"`
	CourseID        int64    `json:"courseid" gorm:"type:bigint;not null;column:courseid"`
	SortOrder       int64    `json:"sortorder" gorm:"type:bigint;not null;default:0;column:sortorder"`
	Name            *string  `json:"name,omitempty" gorm:"type:varchar(255);column:name"`
	EnrolPeriod     *int64   `json:"enrolperiod,omitempty" gorm:"type:bigint;column:enrolperiod;default:0"`
	EnrolStartDate  *int64   `json:"enrolstartdate,omitempty" gorm:"type:bigint;column:enrolstartdate;default:0"`
	EnrolEndDate    *int64   `json:"enrolenddate,omitempty" gorm:"type:bigint;column:enrolenddate;default:0"`
	ExpiryNotify    *int16   `json:"expirynotify,omitempty" gorm:"type:smallint;column:expirynotify;default:0"`
	ExpiryThreshold *int64   `json:"expirythreshold,omitempty" gorm:"type:bigint;column:expirythreshold;default:0"`
	NotifyAll       *int16   `json:"notifyall,omitempty" gorm:"type:smallint;column:notifyall;default:0"`
	Password        *string  `json:"password,omitempty" gorm:"type:varchar(50);column:password"`
	Cost            *string  `json:"cost,omitempty" gorm:"type:varchar(20);column:cost"`
	Currency        *string  `json:"currency,omitempty" gorm:"type:varchar(3);column:currency"`
	RoleID          *int64   `json:"roleid,omitempty" gorm:"type:bigint;column:roleid;default:0"`
	CustomInt1      *int64   `json:"customint1,omitempty" gorm:"type:bigint;column:customint1"`
	CustomInt2      *int64   `json:"customint2,omitempty" gorm:"type:bigint;column:customint2"`
	CustomInt3      *int64   `json:"customint3,omitempty" gorm:"type:bigint;column:customint3"`
	CustomInt4      *int64   `json:"customint4,omitempty" gorm:"type:bigint;column:customint4"`
	CustomInt5      *int64   `json:"customint5,omitempty" gorm:"type:bigint;column:customint5"`
	CustomInt6      *int64   `json:"customint6,omitempty" gorm:"type:bigint;column:customint6"`
	CustomInt7      *int64   `json:"customint7,omitempty" gorm:"type:bigint;column:customint7"`
	CustomInt8      *int64   `json:"customint8,omitempty" gorm:"type:bigint;column:customint8"`
	CustomChar1     *string  `json:"customchar1,omitempty" gorm:"type:varchar(255);column:customchar1"`
	CustomChar2     *string  `json:"customchar2,omitempty" gorm:"type:varchar(255);column:customchar2"`
	CustomChar3     *string  `json:"customchar3,omitempty" gorm:"type:varchar(1333);column:customchar3"`
	CustomDec1      *float64 `json:"customdec1,omitempty" gorm:"type:numeric(12,7);column:customdec1"`
	CustomDec2      *float64 `json:"customdec2,omitempty" gorm:"type:numeric(12,7);column:customdec2"`
	CustomText1     *string  `json:"customtext1,omitempty" gorm:"type:text;column:customtext1"`
	CustomText2     *string  `json:"customtext2,omitempty" gorm:"type:text;column:customtext2"`
	CustomText3     *string  `json:"customtext3,omitempty" gorm:"type:text;column:customtext3"`
	CustomText4     *string  `json:"customtext4,omitempty" gorm:"type:text;column:customtext4"`
	TimeCreated     int64    `json:"timecreated" gorm:"type:bigint;not null;default:0;column:timecreated"`
	TimeModified    int64    `json:"timemodified" gorm:"type:bigint;not null;default:0;column:timemodified"`
}

func (Enrol) TableName() string {
	return "mdl_enrol"
}
