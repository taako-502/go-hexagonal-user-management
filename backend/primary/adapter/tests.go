package primary_adapter

import (
	"go-hexagonal-user-management/config"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func SetupEchoForTest() *echo.Echo {
	e := echo.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}
	return e
}