package model

type Course struct {
	ID                       int64   `json:"id" gorm:"column:id;primaryKey"`
	Category                 int64   `json:"category" gorm:"column:category"`
	SortOrder                int64   `json:"sortorder" gorm:"column:sortorder"`
	FullName                 string  `json:"fullname" gorm:"column:fullname"`
	ShortName                string  `json:"shortname" gorm:"column:shortname"`
	IDNumber                 string  `json:"idnumber" gorm:"column:idnumber"`
	Summary                  *string `json:"summary" gorm:"column:summary"`
	SummaryFormat            int8    `json:"summaryformat" gorm:"column:summaryformat"`
	Format                   string  `json:"format" gorm:"column:format"`
	ShowGrades               int8    `json:"showgrades" gorm:"column:showgrades"`
	NewsItems                int32   `json:"newsitems" gorm:"column:newsitems"`
	StartDate                int64   `json:"startdate" gorm:"column:startdate"`
	EndDate                  int64   `json:"enddate" gorm:"column:enddate"`
	RelativeDatesMode        int8    `json:"relativedatesmode" gorm:"column:relativedatesmode"`
	Marker                   int64   `json:"marker" gorm:"column:marker"`
	MaxBytes                 int64   `json:"maxbytes" gorm:"column:maxbytes"`
	LegacyFiles              int16   `json:"legacyfiles" gorm:"column:legacyfiles"`
	ShowReports              int16   `json:"showreports" gorm:"column:showreports"`
	Visible                  int8    `json:"visible" gorm:"column:visible"`
	VisibleOld               int8    `json:"visibleold" gorm:"column:visibleold"`
	DownloadContent          *int8   `json:"downloadcontent" gorm:"column:downloadcontent"`
	GroupMode                int16   `json:"groupmode" gorm:"column:groupmode"`
	GroupModeForce           int16   `json:"groupmodeforce" gorm:"column:groupmodeforce"`
	DefaultGroupingID        int64   `json:"defaultgroupingid" gorm:"column:defaultgroupingid"`
	Lang                     string  `json:"lang" gorm:"column:lang"`
	CalendarType             string  `json:"calendartype" gorm:"column:calendartype"`
	Theme                    string  `json:"theme" gorm:"column:theme"`
	TimeCreated              int64   `json:"timecreated" gorm:"column:timecreated"`
	TimeModified             int64   `json:"timemodified" gorm:"column:timemodified"`
	Requested                int8    `json:"requested" gorm:"column:requested"`
	EnableCompletion         int8    `json:"enablecompletion" gorm:"column:enablecompletion"`
	CompletionNotify         int8    `json:"completionnotify" gorm:"column:completionnotify"`
	CacheRev                 int64   `json:"cacherev" gorm:"column:cacherev"`
	OriginalCourseID         *int64  `json:"originalcourseid" gorm:"column:originalcourseid"`
	ShowActivityDates        int8    `json:"showactivitydates" gorm:"column:showactivitydates"`
	ShowCompletionConditions *int8   `json:"showcompletionconditions" gorm:"column:showcompletionconditions"`
	PDFExportFont            *string `json:"pdfexportfont" gorm:"column:pdfexportfont"`
}

// TableName sets the insert table name for this struct type
func (Course) TableName() string {
	return "mdl_course"
}
