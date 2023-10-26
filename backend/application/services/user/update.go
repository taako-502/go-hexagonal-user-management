package user_service

import (
	"errors"
	"go-sample-api/application/domain"
	user_secondary_adapter "go-sample-api/secondary/adapter/user"
	secondary_port "go-sample-api/secondary/port"
)

func (u UserService)Update(a secondary_port.UserRepository, user *domain.User) error {
	if err := a.Update(user); err != nil {
		if errors.Is(err, user_secondary_adapter.UserDuplicateError) {
			return UserDuplicateError
		} else {
			return err
		}
	}
	return nil
}