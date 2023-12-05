package secondary_port

import user_model "go-hexagonal-user-management/core/models"

type UserRepository interface {
	Create(user *user_model.User) (*user_model.User, error)
	Update(user *user_model.User) error
	FindAll() ([]user_model.User, error)
	Delete(id int) error
}
