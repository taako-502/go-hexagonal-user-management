package user_service

import (
	secondary_port "go-sample-api/secondary/port"
)

func Delete(a secondary_port.UserRepository, id int) error {
	if err := a.Delete(id); err != nil {
		return err
	}
	return nil
}