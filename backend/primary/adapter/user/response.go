package user_primary_adapter

type UserResponse struct{
	Id int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}