package user_primary_adapter

import (
	"errors"
	"go-sample-api/application/domain"
	user_service "go-sample-api/application/services/user"
	user_primary_port "go-sample-api/primary/port/user"
	user_secondary_adapter "go-sample-api/secondary/adapter/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (db *myDB)Create(c echo.Context) error {
	request := new(user_primary_port.UserRequest)
	c.Bind(request)
	user := &domain.User{
		Username: request.Username,
		Email: request.Email,
	}
	a := user_secondary_adapter.NewUserSecondaryAdapter(db.DB)
	if err := user_service.Create(a, user); err != nil {
		if errors.Is(err, user_service.UserDuplicateError) {
			return echo.NewHTTPError(http.StatusConflict, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.String(http.StatusOK, "OK")
}
