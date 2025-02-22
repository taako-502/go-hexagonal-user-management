package user_service

import (
	"errors"
	"fmt"
	"go-hexagonal-user-management/core/model"
	"go-hexagonal-user-management/secondary/adapter/user_secondary_adapter"
	"go-hexagonal-user-management/secondary/secondary_port"
)

func (u UserService) FindAll(a secondary_port.UserRepository) ([]model.User, error) {
	users, err := a.FindAll()
	if errors.Is(err, user_secondary_adapter.ErrUserNotFound) {
		return nil, fmt.Errorf("user_secondary_adapter.FindAll: %w", ErrUserNotFound)
	} else if err != nil {
		return nil, fmt.Errorf("user_secondary_adapter.FindAll: %w", err)
	}
	return users, nil
}
