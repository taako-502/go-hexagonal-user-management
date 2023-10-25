package user_service

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)
type UserService struct{
	Echo *echo.Echo
	DB *gorm.DB
}