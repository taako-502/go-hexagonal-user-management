package user_service

import (
	user_model "go-hexagonal-user-management/core/models"
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
	secondary_port "go-hexagonal-user-management/secondary/port"

	"github.com/pkg/errors"
)

func (u UserService) FindAll(a secondary_port.UserRepository) ([]user_model.User, error) {
	users, err := a.FindAll()
	if errors.Is(err, user_secondary_adapter.ErrUserNotFound) {
		return nil, errors.Wrap(ErrUserDuplicate, "user_secondary_adapter.FindAll")
	} else if err != nil {
		return nil, errors.Wrap(err, "user_secondary_adapter.FindAll")
	}
	return users, nil
}
