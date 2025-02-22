package user_secondary_adapter

import (
	"fmt"
	user_model "go-hexagonal-user-management/core/models"
	secondary_port "go-hexagonal-user-management/secondary/port"
)

type fakeUserRepository struct {
	insertUser  map[*user_model.User]error
	updateUser  map[*user_model.User]error
	findAllUser []user_model.User
	deleteUser  map[int]error
}

func NewFakeUserRepository() secondary_port.UserRepository {
	return &fakeUserRepository{
		insertUser: map[*user_model.User]error{},
		updateUser: map[*user_model.User]error{},
		findAllUser: []user_model.User{
			{ID: 1, Username: "user1", Email: "user1@example.com"},
			{ID: 2, Username: "user2", Email: "user2@example.com"},
		},
		deleteUser: map[int]error{},
	}
}

func (r *fakeUserRepository) Create(u *user_model.User) (*user_model.User, error) {
	if err := r.insertUser[u]; err != nil {
		return nil, fmt.Errorf("fakeUserRepository.insertUser: %w", err)
	}
	return &user_model.User{
		ID: 999,
	}, nil
}

func (r *fakeUserRepository) Update(u *user_model.User) error {
	if err := r.updateUser[u]; err != nil {
		return fmt.Errorf("fakeUserRepository.updateUser: %w", err)
	}
	return nil
}

func (r *fakeUserRepository) FindAll() ([]user_model.User, error) {
	return r.findAllUser, nil
}

func (r *fakeUserRepository) Delete(id int) error {
	if err := r.deleteUser[id]; err != nil {
		return fmt.Errorf("fakeUserRepository.deleteUser: %w", err)
	}
	return nil
}
