package user_secondary_adapter

import (
	"go-hexagonal-user-management/secondary/secondary_port"

	"gorm.io/gorm"
)

type userSecondaryAdapter struct {
	Db *gorm.DB
}

func NewUserSecondaryAdapter(db *gorm.DB) secondary_port.UserRepository {
	return &userSecondaryAdapter{Db: db}
}
