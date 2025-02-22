package user_primary_adapter

import "github.com/go-playground/validator/v10"

type UserPrimaryAdapter struct {
	validate *validator.Validate
}

func NewUserPrimaryAdapter(validate *validator.Validate) *UserPrimaryAdapter {
	return &UserPrimaryAdapter{validate: validate}
}
