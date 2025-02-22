package primary_port

import (
	"go-hexagonal-user-management/core/model"
	"go-hexagonal-user-management/secondary/secondary_port"
)

type User interface {
	Create(a secondary_port.UserRepository, user *model.User) error
	Update(a secondary_port.UserRepository, user *model.User) error
	FindAll(a secondary_port.UserRepository) ([]model.User, error)
	Delete(a secondary_port.UserRepository, id int) error
}
