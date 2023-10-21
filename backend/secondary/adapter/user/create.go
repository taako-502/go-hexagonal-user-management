package user_secondary_adapter

import (
	"go-sample-api/application/domain"
	user_secondary_port "go-sample-api/secondary/port/user"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type userSecondaryAdapter struct {
  Db *gorm.DB
}

func NewUserSecondaryAdapter(db *gorm.DB) user_secondary_port.UserRepository {
	return &userSecondaryAdapter{Db: db}
}

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

