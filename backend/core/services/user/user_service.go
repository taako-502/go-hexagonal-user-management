package user_service

import (
	"github.com/labstack/echo/v4"
)
type UserService struct{
	Echo *echo.Echo
}

func NewUserService(e *echo.Echo) UserService {
	return UserService{ Echo: e }
}