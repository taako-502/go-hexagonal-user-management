package user_secondary_adapter

import (
	"go-sample-api/application/domain"
	user_secondary_port "go-sample-api/secondary/port/user"

	"gorm.io/gorm"
)

type userSecondaryAdapter struct {
  Db *gorm.DB
}

func NewUserSecondaryAdapter(db *gorm.DB) user_secondary_port.UserRepository {
	return &userSecondaryAdapter{Db: db}
}

func (m *userSecondaryAdapter)Create(user *domain.User) error {
	result := m.Db.Create(user)
  return result.Error
}

