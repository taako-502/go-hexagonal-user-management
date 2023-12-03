package user_secondary_adapter

import (
	user_model "go-hexagonal-user-management/core/models"

	"github.com/go-sql-driver/mysql"
)

func (a *userSecondaryAdapter) Update(user *user_model.User) error {
	if err := a.Db.Model(&user).Updates(user).Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return ErrUserDuplicate
		} else {
			return err
		}
	}
	return nil
}
