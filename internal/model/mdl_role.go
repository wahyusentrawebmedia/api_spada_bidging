package model

type Role struct {
	ID          int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name"`
	ShortName   string `json:"shortname"`
	Description string `json:"description"`
	SortOrder   int64  `json:"sortorder"`
	Archetype   string `json:"archetype"`
}

func (Role) TableName() string {
	return "mdl_role"
}
