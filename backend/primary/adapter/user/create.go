package user_primary_adapter

import (
	"go-sample-api/application/domain"
	user_service "go-sample-api/application/services/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	user := new(domain.User)
	c.Bind(user)
	if err := user_service.Create(user); err != nil {
		return err
	}

	return c.String(http.StatusOK, "Hello, World!")
}