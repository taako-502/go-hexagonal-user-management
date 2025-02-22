package user_secondary_adapter

import (
	"fmt"
	"go-hexagonal-user-management/core/model"
	"go-hexagonal-user-management/secondary/secondary_port"
)

type fakeUserRepository struct {
	insertUser  map[*model.User]error
	updateUser  map[*model.User]error
	findAllUser []model.User
	deleteUser  map[int]error
}

func NewFakeUserRepository() secondary_port.UserRepository {
	return &fakeUserRepository{
		insertUser: map[*model.User]error{},
		updateUser: map[*model.User]error{},
		findAllUser: []model.User{
			{ID: 1, Username: "user1", Email: "user1@example.com"},
			{ID: 2, Username: "user2", Email: "user2@example.com"},
		},
		deleteUser: map[int]error{},
	}
}

func (r *fakeUserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.insertUser[u]; err != nil {
		return nil, fmt.Errorf("fakeUserRepository.insertUser: %w", err)
	}
	return &model.User{
		ID: 999,
	}, nil
}

func (r *fakeUserRepository) Update(u *model.User) error {
	if err := r.updateUser[u]; err != nil {
		return fmt.Errorf("fakeUserRepository.updateUser: %w", err)
	}
	return nil
}

func (r *fakeUserRepository) FindAll() ([]model.User, error) {
	return r.findAllUser, nil
}

func (r *fakeUserRepository) Delete(id int) error {
	if err := r.deleteUser[id]; err != nil {
		return fmt.Errorf("fakeUserRepository.deleteUser: %w", err)
	}
	return nil
}
