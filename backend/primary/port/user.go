package primary_port

import (
	user_model "go-hexagonal-user-management/core/models"
	secondary_port "go-hexagonal-user-management/secondary/port"
)

type User interface {
	Create(a secondary_port.UserRepository, user *user_model.User) error
	Update(a secondary_port.UserRepository, user *user_model.User) error
	FindAll(a secondary_port.UserRepository) ([]user_model.User, error)
	Delete(a secondary_port.UserRepository, id int) error
}
