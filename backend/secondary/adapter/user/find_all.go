package user_secondary_adapter

import (
	user_model "go-hexagonal-user-management/core/models"

	"github.com/pkg/errors"
)

func (a *userSecondaryAdapter) FindAll() ([]user_model.User, error) {
	var users []user_model.User
	result := a.Db.Find(&users)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "gorm.Find")
	}
	if len(users) == 0 {
		return nil, errors.Wrap(ErrUserNotFound, "gorm.Find")
	}
	return users, nil
}
