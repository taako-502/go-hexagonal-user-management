package user_primary_port

type UserRequest struct {
	Username string `json:"username" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
