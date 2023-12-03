package user_secondary_adapter

import (
	user_model "go-hexagonal-user-management/core/models"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func (a *userSecondaryAdapter) Create(user *user_model.User) error {
	if err := a.Db.Create(user).Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return errors.Wrap(ErrUserDuplicate, "gorm.Create")
		} else {
			return errors.Wrap(err, "gorm.Create")
		}
	}
	return nil
}
