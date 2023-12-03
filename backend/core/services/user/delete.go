package user_service

import (
	secondary_port "go-hexagonal-user-management/secondary/port"

	"github.com/pkg/errors"
)

func (u UserService) Delete(a secondary_port.UserRepository, id int) error {
	if err := a.Delete(id); err != nil {
		return errors.Wrap(err, "user_secondary_adapter.Delete")
	}
	return nil
}
