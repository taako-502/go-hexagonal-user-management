package user_service

import (
	"errors"
	"fmt"
	"go-hexagonal-user-management/core/model"
	"go-hexagonal-user-management/secondary/adapter/user_secondary_adapter"
	"go-hexagonal-user-management/secondary/secondary_port"
)

func (u UserService) Update(a secondary_port.UserRepository, user *model.User) error {
	if err := a.Update(user); err != nil {
		if errors.Is(err, user_secondary_adapter.ErrUserDuplicate) {
			return fmt.Errorf("user_secondary_adapter.Update: %w", ErrUserDuplicate)
		} else {
			return fmt.Errorf("user_secondary_adapter.Update: %w", err)
		}
	}
	return nil
}
