package user_service

import (
	"errors"
	"go-hexagonal-user-management/core/domain"
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
	secondary_port "go-hexagonal-user-management/secondary/port"
)

func (u UserService)Create(a secondary_port.UserRepository, user *domain.User) error {
	if err := a.Create(user); err != nil {
		if errors.Is(err, user_secondary_adapter.UserDuplicateError) {
			return UserDuplicateError
		} else {
			return err
		}
	}
	return nil
}