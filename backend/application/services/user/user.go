package user_service

import (
	"go-sample-api/application/domain"
	user_secondary_adapter "go-sample-api/secondary/adapter/user"

	"gorm.io/gorm"
)

func Create(db *gorm.DB, user *domain.User) error {
	u := user_secondary_adapter.NewUserSecondaryAdapter(db)
	return u.Create(user)
}