package model

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MdlUser struct {
	ID                int64   `json:"id" gorm:"column:id;primaryKey"`
	Auth              string  `json:"auth" gorm:"column:auth"`
	Confirmed         int     `json:"confirmed" gorm:"column:confirmed"`
	PolicyAgreed      int     `json:"policy_agreed" gorm:"column:policyagreed"`
	Deleted           int     `json:"deleted" gorm:"column:deleted"`
	Suspended         int     `json:"suspended" gorm:"column:suspended"`
	MnethostID        int64   `json:"mnethost_id" gorm:"column:mnethostid"`
	Username          string  `json:"username" gorm:"column:username"`
	Password          string  `json:"password" gorm:"column:password"`
	IDNumber          string  `json:"id_number" gorm:"column:idnumber"`
	FirstName         string  `json:"first_name" gorm:"column:firstname"`
	LastName          string  `json:"last_name" gorm:"column:lastname"`
	Email             string  `json:"email" gorm:"column:email"`
	EmailStop         int     `json:"email_stop" gorm:"column:emailstop"`
	Phone1            string  `json:"phone1" gorm:"column:phone1"`
	Phone2            string  `json:"phone2" gorm:"column:phone2"`
	Institution       string  `json:"institution" gorm:"column:institution"`
	Department        string  `json:"department" gorm:"column:department"`
	Address           string  `json:"address" gorm:"column:address"`
	City              string  `json:"city" gorm:"column:city"`
	Country           string  `json:"country" gorm:"column:country"`
	Lang              string  `json:"lang" gorm:"column:lang"`
	CalendarType      string  `json:"calendar_type" gorm:"column:calendartype"`
	Theme             string  `json:"theme" gorm:"column:theme"`
	Timezone          string  `json:"timezone" gorm:"column:timezone"`
	FirstAccess       int64   `json:"first_access" gorm:"column:firstaccess"`
	LastAccess        int64   `json:"last_access" gorm:"column:lastaccess"`
	LastLogin         int64   `json:"last_login" gorm:"column:lastlogin"`
	CurrentLogin      int64   `json:"current_login" gorm:"column:currentlogin"`
	LastIP            string  `json:"last_ip" gorm:"column:lastip"`
	Secret            string  `json:"secret" gorm:"column:secret"`
	Picture           int64   `json:"picture" gorm:"column:picture"`
	Description       *string `json:"description,omitempty" gorm:"column:description"`
	DescriptionFormat int     `json:"description_format" gorm:"column:descriptionformat"`
	MailFormat        int     `json:"mail_format" gorm:"column:mailformat"`
	MailDigest        int     `json:"mail_digest" gorm:"column:maildigest"`
	MailDisplay       int     `json:"mail_display" gorm:"column:maildisplay"`
	AutoSubscribe     int     `json:"auto_subscribe" gorm:"column:autosubscribe"`
	TrackForums       int     `json:"track_forums" gorm:"column:trackforums"`
	TimeCreated       int64   `json:"time_created" gorm:"column:timecreated"`
	TimeModified      int64   `json:"time_modified" gorm:"column:timemodified"`
	TrustBitmask      int64   `json:"trust_bitmask" gorm:"column:trustbitmask"`
	ImageAlt          *string `json:"image_alt,omitempty" gorm:"column:imagealt"`
	LastNamePhonetic  *string `json:"last_name_phonetic,omitempty" gorm:"column:lastnamephonetic"`
	FirstNamePhonetic *string `json:"first_name_phonetic,omitempty" gorm:"column:firstnamephonetic"`
	MiddleName        *string `json:"middle_name,omitempty" gorm:"column:middlename"`
	AlternateName     *string `json:"alternate_name,omitempty" gorm:"column:alternatename"`
	MoodleNetProfile  *string `json:"moodlenet_profile,omitempty" gorm:"column:moodlenetprofile"`
}

func (r *MdlUser) TableName() string {
	return "mdl_user"
}
