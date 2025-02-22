package user_secondary_adapter

import (
	"fmt"
	user_model "go-hexagonal-user-management/core/models"

	"github.com/go-sql-driver/mysql"
)

func (a *userSecondaryAdapter) Create(user *user_model.User) (*user_model.User, error) {
	if err := a.Db.Create(user).Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return nil, fmt.Errorf("gorm.Create: %w", ErrUserDuplicate)
		} else {
			return nil, fmt.Errorf("gorm.Create: %w", err)
		}
	}
	return user, nil
}
