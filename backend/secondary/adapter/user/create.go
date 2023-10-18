package user

import (
	"database/sql"
	"go-sample-api/application/user"
)

type MySQLUserRepository struct {
  db *sql.DB
}

func (m *MySQLUserRepository) GetUserByID(id string) (*user.User, error) {
  // MySQLとの通信ロジック
  return &user.User{}, nil
}
