package model

type GradeCategory struct {
	ID                  int64   `json:"id" gorm:"column:id;primaryKey"`
	CourseID            int64   `json:"courseid" gorm:"column:courseid"`
	Parent              *int64  `json:"parent,omitempty" gorm:"column:parent"`
	Depth               int64   `json:"depth" gorm:"column:depth"`
	Path                *string `json:"path,omitempty" gorm:"column:path"`
	FullName            string  `json:"fullname" gorm:"column:fullname"`
	Aggregation         int64   `json:"aggregation" gorm:"column:aggregation"`
	KeepHigh            int64   `json:"keephigh" gorm:"column:keephigh"`
	DropLow             int64   `json:"droplow" gorm:"column:droplow"`
	AggregateOnlyGraded int16   `json:"aggregateonlygraded" gorm:"column:aggregateonlygraded"`
	AggregateOutcomes   int16   `json:"aggregateoutcomes" gorm:"column:aggregateoutcomes"`
	TimeCreated         int64   `json:"timecreated" gorm:"column:timecreated"`
	TimeModified        int64   `json:"timemodified" gorm:"column:timemodified"`
	Hidden              int64   `json:"hidden" gorm:"column:hidden"`
}

func (GradeCategory) TableName() string {
	return "mdl_grade_categories"
}
