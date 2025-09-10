package model

type MdlGroupsMember struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	GroupID   int64  `gorm:"column:groupid;not null;default:0;index:mdl_groumemb_gro_ix" json:"group_id"`
	UserID    int64  `gorm:"column:userid;not null;default:0;index:mdl_groumemb_use_ix;uniqueIndex:mdl_groumemb_usegro_uix,priority:1" json:"user_id"`
	TimeAdded int64  `gorm:"column:timeadded;not null;default:0" json:"time_added"`
	Component string `gorm:"column:component;type:varchar(100);not null;default:''" json:"component"`
	ItemID    int64  `gorm:"column:itemid;not null;default:0" json:"item_id"`
}

// TableName overrides the table name used by GORM.
func (MdlGroupsMember) TableName() string {
	return "mdl_groups_members"
}
