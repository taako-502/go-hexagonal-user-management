package secondary_port

import (
	"go-sample-api/application/domain"
)

type UserRepository interface {
  Create(user *domain.User) error
	Update(user *domain.User) error
	FindAll() ([]domain.User, error)
	Delete(id int) error
}

