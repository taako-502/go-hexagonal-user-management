package user_secondary_adapter

import (
	"fmt"
	"go-hexagonal-user-management/core/model"

	"github.com/go-sql-driver/mysql"
)

func (a *userSecondaryAdapter) Update(user *model.User) error {
	if err := a.Db.Model(&user).Updates(user).Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return fmt.Errorf("gorm.Updates: %w", ErrUserDuplicate)
		} else {
			return fmt.Errorf("gorm.Updates: %w", err)
		}
	}
	return nil
}
