package user_service

import (
	"errors"
	user_model "go-hexagonal-user-management/core/models"
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
	secondary_port "go-hexagonal-user-management/secondary/port"
)

func (u UserService) FindAll(a secondary_port.UserRepository) ([]user_model.User, error) {
	users, err := a.FindAll()
	if errors.Is(err, user_secondary_adapter.ErrUserNotFound) {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return users, nil
}
