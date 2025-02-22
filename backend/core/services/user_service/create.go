package user_service

import (
	"fmt"
	"go-hexagonal-user-management/core/model"
	"go-hexagonal-user-management/secondary/adapter/user_secondary_adapter"
	"go-hexagonal-user-management/secondary/secondary_port"

	"errors"
)

func (u UserService) Create(a secondary_port.UserRepository, user *model.User) (*model.User, error) {
	user, err := a.Create(user)
	if err != nil {
		if errors.Is(err, user_secondary_adapter.ErrUserDuplicate) {
			return nil, fmt.Errorf("user_secondary_adapter.Create: %w", ErrUserDuplicate)
		} else {
			return nil, fmt.Errorf("user_secondary_adapter.Create: %w", err)
		}
	}
	return user, nil
}
