package model

type UserSyncRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserSyncResponse struct {
	Action   bool   `json:"action"`
	Username string `json:"username"`
	Password string `json:"password"`
	Pesan    string `json:"pesan"`
	IdSpada  int64  `json:"id_spada"`
}
