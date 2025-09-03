package model

type MdlCourseCategory struct {
	ID                int64   `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name              string  `gorm:"type:varchar(255);not null;default:'';column:name" json:"name"`
	IDNumber          *string `gorm:"type:varchar(100);column:idnumber" json:"idnumber,omitempty"`
	Description       *string `gorm:"type:text;column:description" json:"description,omitempty"`
	DescriptionFormat int16   `gorm:"type:int2;not null;default:0;column:descriptionformat" json:"description_format"`
	Parent            int64   `gorm:"type:int8;not null;default:0;column:parent" json:"parent"`
	SortOrder         int64   `gorm:"type:int8;not null;default:0;column:sortorder" json:"sort_order"`
	CourseCount       int64   `gorm:"type:int8;not null;default:0;column:coursecount" json:"course_count"`
	Visible           int16   `gorm:"type:int2;not null;default:1;column:visible" json:"visible"`
	VisibleOld        int16   `gorm:"type:int2;not null;default:1;column:visibleold" json:"visible_old"`
	TimeModified      int64   `gorm:"type:int8;not null;default:0;column:timemodified" json:"time_modified"`
	Depth             int64   `gorm:"type:int8;not null;default:0;column:depth" json:"depth"`
	Path              string  `gorm:"type:varchar(255);not null;default:'';column:path" json:"path"`
	Theme             *string `gorm:"type:varchar(50);column:theme" json:"theme,omitempty"`
}

// TableName overrides the table name used by GORM
func (MdlCourseCategory) TableName() string {
	return "mdl_course_categories"
}
