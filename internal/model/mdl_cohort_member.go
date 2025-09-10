package model

type CohortMember struct {
	ID        int64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CohortID  int64 `gorm:"column:cohortid;not null;index:idx_cohort_user,unique;index" json:"cohortid"`
	UserID    int64 `gorm:"column:userid;not null;index:idx_cohort_user,unique;index" json:"userid"`
	TimeAdded int64 `gorm:"column:timeadded;not null" json:"timeadded"`
}

// TableName overrides the table name used by GORM
func (CohortMember) TableName() string {
	return "mdl_cohort_members"
}
