package user_secondary_adapter

import (
	"go-sample-api/application/domain"

	"gorm.io/gorm"
)

type userSecondaryAdapter struct {
  Db *gorm.DB
}

func NewUserSecondaryAdapter(db *gorm.DB) *userSecondaryAdapter {
	return &userSecondaryAdapter{Db: db}
}

func (m *userSecondaryAdapter)Create(user *domain.User) error {
	result := m.Db.Create(user)
  return result.Error
}

