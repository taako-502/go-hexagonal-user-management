package user_primary_adapter

import (
	"go-hexagonal-user-management/core/domain"
	user_service "go-hexagonal-user-management/core/services/user"
	secondary_port "go-hexagonal-user-management/secondary/port"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Update(u user_service.UserService, a secondary_port.UserRepository) *echo.Echo {
	u.Echo.PATCH("/user/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "ID must be an integer")
		}
		request := new(UserRequest)
		if err := c.Bind(request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// validation（https://echo.labstack.com/docs/request#validate-data）
		if err := c.Validate(request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "validation error")
		}
		user := &domain.User{
			Id:       id,
			Username: request.Username,
			Email:    request.Email,
		}
		if err := u.Update(a, user); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, "OK")
	})

	return u.Echo
}
