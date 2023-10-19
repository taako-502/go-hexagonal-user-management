package user_service

import (
	"go-sample-api/application/domain"
	user_secondary_adapter "go-sample-api/secondary/adapter/user"
)

func Create(user *domain.User) error {
	u := user_secondary_adapter.UserSecondaryAdapter{}
	return u.Create(user)
}