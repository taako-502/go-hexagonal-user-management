package user_service

import (
	"errors"
	"go-hexagonal-user-management/core/domain"
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
	secondary_port "go-hexagonal-user-management/secondary/port"
)

func (u UserService) FindAll(a secondary_port.UserRepository) ([]domain.User, error) {
	users, err := a.FindAll()
	if errors.Is(err, user_secondary_adapter.ErrUserNotFound) {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return users, nil
}
