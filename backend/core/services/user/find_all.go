package user_service

import (
	"errors"
	"go-hexagonal-user-management/core/domain"
	secondary_port "go-hexagonal-user-management/secondary/port"
)

func (u UserService)FindAll(a secondary_port.UserRepository) ([]domain.User, error) {
	users, err := a.FindAll()
	if errors.Is(err, UserNotFoundError) {
		return nil, UserNotFoundError
	} else if err != nil {
		return nil, err
	}
	return users, nil
}