package user_secondary_adapter

import (
	"go-hexagonal-user-management/core/domain"

	"github.com/go-sql-driver/mysql"
)

func (a *userSecondaryAdapter)Create(user *domain.User) error {
	if err := a.Db.Create(user); err != nil {
		mysqlErr, ok := err.Error.(*mysql.MySQLError)
		if ok && mysqlErr.Number == 1062 {
			return UserDuplicateError
		} else {
			return err.Error
		}
	}
  return nil
}

