package user

import "go-sample-api/application/user"

type UserRepository interface {
  GetUserByID(id string) (*user.User, error)
}
