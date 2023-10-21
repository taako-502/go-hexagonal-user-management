package user_secondary_adapter

import (
	"go-sample-api/application/domain"
	user_secondary_port "go-sample-api/secondary/port/user"
)

type fakeUserRepository struct {
	insertUser map[*domain.User]error
}

func NewFakeUserRepository() user_secondary_port.UserRepository {
	return &fakeUserRepository{insertUser: map[*domain.User]error{}}
}

func (r *fakeUserRepository) Create(u *domain.User) error {
	if err := r.insertUser[u]; err != nil {
		return err
	}
	return nil
}
