package user_secondary_adapter

import (
	"go-hexagonal-user-management/core/domain"
)

func (a *userSecondaryAdapter)Update(user *domain.User) error {
	if err := a.Db.Model(&user).Updates(user).Error; err != nil {
		return err
	}
  return nil
}

