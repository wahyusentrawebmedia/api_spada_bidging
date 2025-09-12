package model

type RoleCapability struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement"`
	ContextID    int64  `gorm:"column:contextid;not null;default:0"`
	RoleID       int64  `gorm:"column:roleid;not null;default:0"`
	Capability   string `gorm:"column:capability;size:255;not null;default:''"`
	Permission   int64  `gorm:"column:permission;not null;default:0"`
	TimeModified int64  `gorm:"column:timemodified;not null;default:0"`
	ModifierID   int64  `gorm:"column:modifierid;not null;default:0"`
}

// TableName overrides the table name used by GORM.
func (RoleCapability) TableName() string {
	return "mdl_role_capabilities"
}
