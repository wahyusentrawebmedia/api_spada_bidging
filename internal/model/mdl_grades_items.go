package model

type MdlGradeItems struct {
	ID               int64   `json:"id" gorm:"column:id"`
	CourseID         *int64  `json:"courseid" gorm:"column:courseid"`
	CategoryID       *int64  `json:"categoryid" gorm:"column:categoryid"`
	ItemName         *string `json:"itemname" gorm:"column:itemname"`
	ItemType         string  `json:"itemtype" gorm:"column:itemtype"`
	ItemModule       *string `json:"itemmodule" gorm:"column:itemmodule"`
	ItemInstance     *int64  `json:"iteminstance" gorm:"column:iteminstance"`
	ItemNumber       *int64  `json:"itemnumber" gorm:"column:itemnumber"`
	ItemInfo         *string `json:"iteminfo" gorm:"column:iteminfo"`
	IDNumber         *string `json:"idnumber" gorm:"column:idnumber"`
	Calculation      *string `json:"calculation" gorm:"column:calculation"`
	GradeType        int16   `json:"gradetype" gorm:"column:gradetype"`
	GradeMax         float64 `json:"grademax" gorm:"column:grademax"`
	GradeMin         float64 `json:"grademin" gorm:"column:grademin"`
	ScaleID          *int64  `json:"scaleid" gorm:"column:scaleid"`
	OutcomeID        *int64  `json:"outcomeid" gorm:"column:outcomeid"`
	GradePass        float64 `json:"gradepass" gorm:"column:gradepass"`
	MultFactor       float64 `json:"multfactor" gorm:"column:multfactor"`
	PlusFactor       float64 `json:"plusfactor" gorm:"column:plusfactor"`
	AggregationCoef  float64 `json:"aggregationcoef" gorm:"column:aggregationcoef"`
	AggregationCoef2 float64 `json:"aggregationcoef2" gorm:"column:aggregationcoef2"`
	SortOrder        int64   `json:"sortorder" gorm:"column:sortorder"`
	Display          int64   `json:"display" gorm:"column:display"`
	Decimals         *int16  `json:"decimals" gorm:"column:decimals"`
	Hidden           int64   `json:"hidden" gorm:"column:hidden"`
	Locked           int64   `json:"locked" gorm:"column:locked"`
	LockTime         int64   `json:"locktime" gorm:"column:locktime"`
	NeedsUpdate      int64   `json:"needsupdate" gorm:"column:needsupdate"`
	WeightOverride   int16   `json:"weightoverride" gorm:"column:weightoverride"`
	TimeCreated      *int64  `json:"timecreated" gorm:"column:timecreated"`
	TimeModified     *int64  `json:"timemodified" gorm:"column:timemodified"`
}

// TableName sets the insert table name for this struct type
func (MdlGradeItems) TableName() string {
	return "mdl_grade_items"
}
