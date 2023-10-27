package primary_port

import (
	"go-hexagonal-user-management/core/domain"
	secondary_port "go-hexagonal-user-management/secondary/port"
)

type User interface {
	Create(a secondary_port.UserRepository, user *domain.User) error
	Update(a secondary_port.UserRepository, user *domain.User) error
	FindAll(a secondary_port.UserRepository) ([]domain.User, error)
	Delete(a secondary_port.UserRepository, id int) error
}