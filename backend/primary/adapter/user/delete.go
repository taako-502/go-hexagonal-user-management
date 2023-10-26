package user_primary_adapter

import (
	user_service "go-sample-api/application/services/user"
	secondary_port "go-sample-api/secondary/port"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Delete(u user_service.UserService, a secondary_port.UserRepository) *echo.Echo {
	u.Echo.DELETE("/user/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "ID must be an integer")
		}
		if err := u.Delete(a , id); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, "OK")
	})
	return u.Echo
}