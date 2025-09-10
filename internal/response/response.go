package response

type DefaultResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}

type UserChangeEmailRequest struct {
	Username string `json:"username" validate:"required"`
	NewEmail string `json:"new-email" validate:"required,email"`
}
