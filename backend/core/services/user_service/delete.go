package user_service

import (
	"fmt"
	"go-hexagonal-user-management/secondary/secondary_port"
)

func (u UserService) Delete(a secondary_port.UserRepository, id int) error {
	if err := a.Delete(id); err != nil {
		return fmt.Errorf("user_secondary_adapter.Delete: %w", err)
	}
	return nil
}
