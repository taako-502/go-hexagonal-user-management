package user_service

import (
	"errors"
	"go-sample-api/application/domain"
	secondary_port "go-sample-api/secondary/port"
)

func FindAll(a secondary_port.UserRepository) ([]domain.User, error) {
	users, err := a.FindAll()
	if errors.Is(err, UserNotFoundError) {
		return nil, UserNotFoundError
	} else if err != nil {
		return nil, err
	}
	return users, nil
}