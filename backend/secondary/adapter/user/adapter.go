package user_secondary_adapter

import (
	secondary_port "go-sample-api/secondary/port"

	"gorm.io/gorm"
)
type userSecondaryAdapter struct {
  Db *gorm.DB
}

func NewUserSecondaryAdapter(db *gorm.DB) secondary_port.UserRepository {
	return &userSecondaryAdapter{Db: db}
}
