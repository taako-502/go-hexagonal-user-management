package primary_port

import (
	"go-sample-api/application/domain"
	secondary_port "go-sample-api/secondary/port"
)

type User interface {
	Create(a secondary_port.UserRepository, user *domain.User) error
	Update(a secondary_port.UserRepository, user *domain.User) error
	FindAll(a secondary_port.UserRepository) ([]domain.User, error)
	Delete(a secondary_port.UserRepository, id int) error
}