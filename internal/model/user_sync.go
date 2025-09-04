package model

type UserSyncRequest struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

type UserSyncResponse struct {
	Action   bool   `json:"action"`
	Username string `json:"username"`
	Password string `json:"password"`
	Pesan    string `json:"pesan"`
	IdSpada  int64  `json:"id_spada"`
}
