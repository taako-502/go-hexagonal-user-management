package user_service

import (
	"go-sample-api/application/domain"
	user_secondary_port "go-sample-api/secondary/port/user"
)

func Create(a user_secondary_port.UserRepository, user *domain.User) error {
	return a.Create(user)
}