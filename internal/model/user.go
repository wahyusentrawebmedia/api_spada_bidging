package model

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MdlUser struct {
	ID                int64   `json:"id" db:"id"`
	Auth              string  `json:"auth" db:"auth"`
	Confirmed         int     `json:"confirmed" db:"confirmed"`
	PolicyAgreed      int     `json:"policy_agreed" db:"policyagreed"`
	Deleted           int     `json:"deleted" db:"deleted"`
	Suspended         int     `json:"suspended" db:"suspended"`
	MnethostID        int64   `json:"mnethost_id" db:"mnethostid"`
	Username          string  `json:"username" db:"username"`
	Password          string  `json:"password" db:"password"`
	IDNumber          string  `json:"id_number" db:"idnumber"`
	FirstName         string  `json:"first_name" db:"firstname"`
	LastName          string  `json:"last_name" db:"lastname"`
	Email             string  `json:"email" db:"email"`
	EmailStop         int     `json:"email_stop" db:"emailstop"`
	Phone1            string  `json:"phone1" db:"phone1"`
	Phone2            string  `json:"phone2" db:"phone2"`
	Institution       string  `json:"institution" db:"institution"`
	Department        string  `json:"department" db:"department"`
	Address           string  `json:"address" db:"address"`
	City              string  `json:"city" db:"city"`
	Country           string  `json:"country" db:"country"`
	Lang              string  `json:"lang" db:"lang"`
	CalendarType      string  `json:"calendar_type" db:"calendartype"`
	Theme             string  `json:"theme" db:"theme"`
	Timezone          string  `json:"timezone" db:"timezone"`
	FirstAccess       int64   `json:"first_access" db:"firstaccess"`
	LastAccess        int64   `json:"last_access" db:"lastaccess"`
	LastLogin         int64   `json:"last_login" db:"lastlogin"`
	CurrentLogin      int64   `json:"current_login" db:"currentlogin"`
	LastIP            string  `json:"last_ip" db:"lastip"`
	Secret            string  `json:"secret" db:"secret"`
	Picture           int64   `json:"picture" db:"picture"`
	Description       *string `json:"description,omitempty" db:"description"`
	DescriptionFormat int     `json:"description_format" db:"descriptionformat"`
	MailFormat        int     `json:"mail_format" db:"mailformat"`
	MailDigest        int     `json:"mail_digest" db:"maildigest"`
	MailDisplay       int     `json:"mail_display" db:"maildisplay"`
	AutoSubscribe     int     `json:"auto_subscribe" db:"autosubscribe"`
	TrackForums       int     `json:"track_forums" db:"trackforums"`
	TimeCreated       int64   `json:"time_created" db:"timecreated"`
	TimeModified      int64   `json:"time_modified" db:"timemodified"`
	TrustBitmask      int64   `json:"trust_bitmask" db:"trustbitmask"`
	ImageAlt          *string `json:"image_alt,omitempty" db:"imagealt"`
	LastNamePhonetic  *string `json:"last_name_phonetic,omitempty" db:"lastnamephonetic"`
	FirstNamePhonetic *string `json:"first_name_phonetic,omitempty" db:"firstnamephonetic"`
	MiddleName        *string `json:"middle_name,omitempty" db:"middlename"`
	AlternateName     *string `json:"alternate_name,omitempty" db:"alternatename"`
	MoodleNetProfile  *string `json:"moodlenet_profile,omitempty" db:"moodlenetprofile"`
}

func (r *MdlUser) TableName() string {
	return "mdl_user"
}
