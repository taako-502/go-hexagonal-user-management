package user_secondary_adapter

import (
	"fmt"
	"go-hexagonal-user-management/core/model"
)

func (a *userSecondaryAdapter) FindAll() ([]model.User, error) {
	var users []model.User
	result := a.Db.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("gorm.Find: %w", result.Error)
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("gorm.Find: %w", ErrUserNotFound)
	}
	return users, nil
}
