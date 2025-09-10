package model

type RoleAssignment struct {
	ID           int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	RoleID       int64  `json:"roleid" gorm:"column:roleid;not null;default:0"`
	ContextID    int64  `json:"contextid" gorm:"column:contextid;not null;default:0"`
	UserID       int64  `json:"userid" gorm:"column:userid;not null;default:0"`
	TimeModified int64  `json:"timemodified" gorm:"column:timemodified;not null;default:0"`
	ModifierID   int64  `json:"modifierid" gorm:"column:modifierid;not null;default:0"`
	Component    string `json:"component" gorm:"column:component;size:100;not null;default:''"`
	ItemID       int64  `json:"itemid" gorm:"column:itemid;not null;default:0"`
	SortOrder    int64  `json:"sortorder" gorm:"column:sortorder;not null;default:0"`
}

// TableName overrides the table name used by GORM.
func (RoleAssignment) TableName() string {
	return "mdl_role_assignments"
}
