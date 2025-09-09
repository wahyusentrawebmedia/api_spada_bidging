package model

type Course struct {
	ID                       int64   `json:"id" db:"id"`
	Category                 int64   `json:"category" db:"category"`
	SortOrder                int64   `json:"sortorder" db:"sortorder"`
	FullName                 string  `json:"fullname" db:"fullname"`
	ShortName                string  `json:"shortname" db:"shortname"`
	IDNumber                 string  `json:"idnumber" db:"idnumber"`
	Summary                  *string `json:"summary" db:"summary"`
	SummaryFormat            int8    `json:"summaryformat" db:"summaryformat"`
	Format                   string  `json:"format" db:"format"`
	ShowGrades               int8    `json:"showgrades" db:"showgrades"`
	NewsItems                int32   `json:"newsitems" db:"newsitems"`
	StartDate                int64   `json:"startdate" db:"startdate"`
	EndDate                  int64   `json:"enddate" db:"enddate"`
	RelativeDatesMode        int8    `json:"relativedatesmode" db:"relativedatesmode"`
	Marker                   int64   `json:"marker" db:"marker"`
	MaxBytes                 int64   `json:"maxbytes" db:"maxbytes"`
	LegacyFiles              int16   `json:"legacyfiles" db:"legacyfiles"`
	ShowReports              int16   `json:"showreports" db:"showreports"`
	Visible                  int8    `json:"visible" db:"visible"`
	VisibleOld               int8    `json:"visibleold" db:"visibleold"`
	DownloadContent          *int8   `json:"downloadcontent" db:"downloadcontent"`
	GroupMode                int16   `json:"groupmode" db:"groupmode"`
	GroupModeForce           int16   `json:"groupmodeforce" db:"groupmodeforce"`
	DefaultGroupingID        int64   `json:"defaultgroupingid" db:"defaultgroupingid"`
	Lang                     string  `json:"lang" db:"lang"`
	CalendarType             string  `json:"calendartype" db:"calendartype"`
	Theme                    string  `json:"theme" db:"theme"`
	TimeCreated              int64   `json:"timecreated" db:"timecreated"`
	TimeModified             int64   `json:"timemodified" db:"timemodified"`
	Requested                int8    `json:"requested" db:"requested"`
	EnableCompletion         int8    `json:"enablecompletion" db:"enablecompletion"`
	CompletionNotify         int8    `json:"completionnotify" db:"completionnotify"`
	CacheRev                 int64   `json:"cacherev" db:"cacherev"`
	OriginalCourseID         *int64  `json:"originalcourseid" db:"originalcourseid"`
	ShowActivityDates        int8    `json:"showactivitydates" db:"showactivitydates"`
	ShowCompletionConditions *int8   `json:"showcompletionconditions" db:"showcompletionconditions"`
	PDFExportFont            *string `json:"pdfexportfont" db:"pdfexportfont"`
}

// TableName sets the insert table name for this struct type
func (c *Course) TableName() string {
	return "mdl_course"
}
