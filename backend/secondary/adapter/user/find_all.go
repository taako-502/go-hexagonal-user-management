package user_secondary_adapter

import (
	"fmt"
	user_model "go-hexagonal-user-management/core/models"
)

func (a *userSecondaryAdapter) FindAll() ([]user_model.User, error) {
	var users []user_model.User
	result := a.Db.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("gorm.Find: %w", result.Error)
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("gorm.Find: %w", ErrUserNotFound)
	}
	return users, nil
}
