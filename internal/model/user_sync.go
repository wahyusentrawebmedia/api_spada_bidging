package model

type DosenMahasiwaSyncRequest struct {
	Mahasiswa []UserSyncRequest `json:"mahasiswa" validate:"dive"`
	Dosen     []UserSyncRequest `json:"dosen" validate:"dive"`
}

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

type UserChangePasswordRequest struct {
	Username    string `json:"username" validate:"required"`
	OldPassword string `json:"password_lama" validate:"required"`
	NewPassword string `json:"password_baru" validate:"required"`
}
