package user_secondary_adapter

import (
	"fmt"
	user_model "go-hexagonal-user-management/core/models"
)

func (a *userSecondaryAdapter) Delete(id int) error {
	user := &user_model.User{Id: id}
	if err := a.Db.Delete(user); err != nil {
		return fmt.Errorf("gorm.Delete: %w", err.Error)
	}
	return nil
}
