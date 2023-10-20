package user_service

import (
	"errors"
	"go-sample-api/application/domain"
	user_secondary_adapter "go-sample-api/secondary/adapter/user"
	user_secondary_port "go-sample-api/secondary/port/user"
)

func Create(a user_secondary_port.UserRepository, user *domain.User) error {
	if err := a.Create(user); err != nil {
		if errors.Is(err, user_secondary_adapter.UserDuplicateError) {
			return UserDuplicateError
		} else {
			return err
		}
	}
	return nil
}