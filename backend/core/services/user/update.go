package user_service

import (
	"errors"
	"fmt"
	user_model "go-hexagonal-user-management/core/models"
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
	secondary_port "go-hexagonal-user-management/secondary/port"
)

func (u UserService) Update(a secondary_port.UserRepository, user *user_model.User) error {
	if err := a.Update(user); err != nil {
		if errors.Is(err, user_secondary_adapter.ErrUserDuplicate) {
			return fmt.Errorf("user_secondary_adapter.Update: %w", ErrUserDuplicate)
		} else {
			return fmt.Errorf("user_secondary_adapter.Update: %w", err)
		}
	}
	return nil
}
