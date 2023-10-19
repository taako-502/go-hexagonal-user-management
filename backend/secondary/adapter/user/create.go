package user_secondary_adapter

import (
	"database/sql"
	"go-sample-api/application/domain"
)

type UserSecondaryAdapter struct {
  db *sql.DB
}

func (m *UserSecondaryAdapter)Create(user *domain.User) error {
  return nil
}

