package user_secondary_adapter

import (
	"fmt"
	"go-hexagonal-user-management/core/model"
)

func (a *userSecondaryAdapter) Delete(id int) error {
	user := &model.User{ID: id}
	if err := a.Db.Delete(user).Error; err != nil {
		return fmt.Errorf("gorm.Delete: %w", err)
	}
	return nil
}
