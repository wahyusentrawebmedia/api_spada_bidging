package model

type GradeGrade struct {
	ID                int64    `json:"id" gorm:"column:id;primaryKey"`
	ItemID            int64    `json:"itemid" gorm:"column:itemid"`
	UserID            int64    `json:"userid" gorm:"column:userid"`
	RawGrade          *float64 `json:"rawgrade,omitempty" gorm:"column:rawgrade"`
	RawGradeMax       float64  `json:"rawgrademax" gorm:"column:rawgrademax"`
	RawGradeMin       float64  `json:"rawgrademin" gorm:"column:rawgrademin"`
	RawScaleID        *int64   `json:"rawscaleid,omitempty" gorm:"column:rawscaleid"`
	UserModified      *int64   `json:"usermodified,omitempty" gorm:"column:usermodified"`
	FinalGrade        *float64 `json:"finalgrade,omitempty" gorm:"column:finalgrade"`
	Hidden            int64    `json:"hidden" gorm:"column:hidden"`
	Locked            int64    `json:"locked" gorm:"column:locked"`
	LockTime          int64    `json:"locktime" gorm:"column:locktime"`
	Exported          int64    `json:"exported" gorm:"column:exported"`
	Overridden        int64    `json:"overridden" gorm:"column:overridden"`
	Excluded          int64    `json:"excluded" gorm:"column:excluded"`
	Feedback          *string  `json:"feedback,omitempty" gorm:"column:feedback"`
	FeedbackFormat    int64    `json:"feedbackformat" gorm:"column:feedbackformat"`
	Information       *string  `json:"information,omitempty" gorm:"column:information"`
	InformationFormat int64    `json:"informationformat" gorm:"column:informationformat"`
	TimeCreated       *int64   `json:"timecreated,omitempty" gorm:"column:timecreated"`
	TimeModified      *int64   `json:"timemodified,omitempty" gorm:"column:timemodified"`
	AggregationStatus string   `json:"aggregationstatus" gorm:"column:aggregationstatus"`
	AggregationWeight *float64 `json:"aggregationweight,omitempty" gorm:"column:aggregationweight"`
}

// TableName returns the table name for GradeGrade struct
func (GradeGrade) TableName() string {
	return "mdl_grade_grades"
}
