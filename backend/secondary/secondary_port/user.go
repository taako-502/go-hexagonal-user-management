package secondary_port

import (
	"go-hexagonal-user-management/core/model"
)

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) error
	FindAll() ([]model.User, error)
	Delete(id int) error
}
