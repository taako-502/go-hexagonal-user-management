package user_secondary_adapter

import (
	"go-sample-api/application/domain"
	secondary_port "go-sample-api/secondary/port"
)

type fakeUserRepository struct {
	insertUser map[*domain.User]error
	updateUser map[*domain.User]error
	findAllUser []domain.User
	deleteUser map[int]error
}

func NewFakeUserRepository() secondary_port.UserRepository {
	return &fakeUserRepository{
		insertUser: map[*domain.User]error{},
		updateUser: map[*domain.User]error{},
		findAllUser: []domain.User{},
		deleteUser: map[int]error{},
	}
}

func (r *fakeUserRepository) Create(u *domain.User) error {
	if err := r.insertUser[u]; err != nil {
		return err
	}
	return nil
}

func (r *fakeUserRepository) Update(u *domain.User) error {
	if err := r.updateUser[u]; err != nil {
		return err
	}
	return nil
}

func (r *fakeUserRepository) FindAll() ([]domain.User, error) {
	users := []domain.User{
		{Username: "test", Email: "test@test.com"},
	}
	return users, nil
}

func (r *fakeUserRepository) Delete(id int) error {
	if err := r.deleteUser[id]; err != nil {
		return err
	}
	return nil
}

