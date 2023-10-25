package user_service

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)
type UserService struct{
	Echo *echo.Echo
}

func NewUserService(e *echo.Echo, db *gorm.DB) *UserService {
	return &UserService{
		Echo: e,
	}
}