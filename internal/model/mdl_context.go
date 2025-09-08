package model

type MdlContext struct {
	ID           int64   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ContextLevel int64   `gorm:"column:contextlevel;not null;default:0" json:"contextlevel"`
	InstanceID   int64   `gorm:"column:instanceid;not null;default:0" json:"instanceid"`
	Path         *string `gorm:"column:path;size:255" json:"path,omitempty"`
	Depth        int8    `gorm:"column:depth;not null;default:0" json:"depth"`
	Locked       int8    `gorm:"column:locked;not null;default:0" json:"locked"`
}

// TableName overrides the table name used by GORM.
func (MdlContext) TableName() string {
	return "mdl_context"
}
