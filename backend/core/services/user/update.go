package user_service

import (
	"go-hexagonal-user-management/core/domain"
	secondary_port "go-hexagonal-user-management/secondary/port"
)

func (u UserService) Update(a secondary_port.UserRepository, user *domain.User) error {
	if err := a.Update(user); err != nil {
		return err
	}
	return nil
}
