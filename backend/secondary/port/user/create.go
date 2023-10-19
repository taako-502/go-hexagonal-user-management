package user

import (
	"go-sample-api/application/domain"
)

type UserRepository interface {
  Create(user *domain.User) (domain.User, error)
}
