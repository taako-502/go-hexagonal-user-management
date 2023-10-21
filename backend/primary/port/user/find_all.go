package user_primary_port

type UserResponse struct{
	Id int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}