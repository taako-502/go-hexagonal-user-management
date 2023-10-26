package user_secondary_adapter

import (
	"go-sample-api/application/domain"

	"github.com/go-sql-driver/mysql"
)

func (a *userSecondaryAdapter)Update(user *domain.User) error {
	if err := a.Db.Model(&user).Updates(user).Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return UserDuplicateError
			}
		}
		return err
	}
  return nil
}
