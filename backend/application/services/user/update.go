package user_service

import (
	"go-sample-api/application/domain"
	secondary_port "go-sample-api/secondary/port"
)

func (u UserService)Update(a secondary_port.UserRepository, user *domain.User) error {
	if err := a.Update(user); err != nil {
		return err
	}
	return nil
}