package user_primary_adapter

import (
	"errors"
	"go-sample-api/application/domain"
	user_service "go-sample-api/application/services/user"
	secondary_port "go-sample-api/secondary/port"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Update(u user_service.UserService, a secondary_port.UserRepository)  *echo.Echo {
	u.Echo.PATCH("/user", func(c echo.Context) error {
		request := new(UserRequest)
		if err := c.Bind(request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// validation（https://echo.labstack.com/docs/request#validate-data）
		if err := c.Validate(request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "validation error")
		}
		user := &domain.User{
			Username: request.Username,
			Email: request.Email,
		}
		if err := u.Update(a, user); err != nil {
			if errors.Is(err, user_service.UserDuplicateError) {
				return echo.NewHTTPError(http.StatusConflict, err.Error())
			} else {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}
		return c.String(http.StatusOK, "OK")
	})

	return u.Echo
}
