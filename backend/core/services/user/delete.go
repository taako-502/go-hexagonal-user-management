package user_service

import (
	"fmt"
	secondary_port "go-hexagonal-user-management/secondary/port"
)

func (u UserService) Delete(a secondary_port.UserRepository, id int) error {
	if err := a.Delete(id); err != nil {
		return fmt.Errorf("user_secondary_adapter.Delete: %w", err)
	}
	return nil
}
