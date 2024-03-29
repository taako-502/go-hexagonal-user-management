package user_primary_adapter

import (
	"errors"
	user_model "go-hexagonal-user-management/core/models"
	user_service "go-hexagonal-user-management/core/services/user"
	secondary_port "go-hexagonal-user-management/secondary/port"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Update(u user_service.UserService, a secondary_port.UserRepository) *echo.Echo {
	u.Echo.PUT("/user/:id", func(c echo.Context) error {
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
		user, err := user_model.UpdateUser(id, request.Username, request.Email)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := u.Update(a, user); err != nil {
			if errors.Is(err, user_service.ErrUserDuplicate) {
				return echo.NewHTTPError(http.StatusConflict, err.Error())
			} else {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}
		return c.String(http.StatusOK, "OK")
	})

	return u.Echo
}
