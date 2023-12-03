package user_secondary_adapter

import (
	"go-hexagonal-user-management/core/domain"

	"github.com/go-sql-driver/mysql"
)

func (a *userSecondaryAdapter) Create(user *domain.User) error {
	if err := a.Db.Create(user).Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return ErrUserDuplicate
		} else {
			return err
		}
	}
	return nil
}
