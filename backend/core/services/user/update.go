package user_service

import (
	user_model "go-hexagonal-user-management/core/models"
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
	secondary_port "go-hexagonal-user-management/secondary/port"

	"github.com/pkg/errors"
)

func (u UserService) Update(a secondary_port.UserRepository, user *user_model.User) error {
	if err := a.Update(user); err != nil {
		if errors.Is(err, user_secondary_adapter.ErrUserDuplicate) {
			return errors.Wrap(ErrUserDuplicate, "user_secondary_adapter.Update")
		} else {
			return errors.Wrap(err, "user_secondary_adapter.Update")
		}
	}
	return nil
}
