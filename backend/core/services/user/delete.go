package user_service

import (
	secondary_port "go-hexagonal-user-management/secondary/port"
)

func (u UserService) Delete(a secondary_port.UserRepository, id int) error {
	if err := a.Delete(id); err != nil {
		return err
	}
	return nil
}
