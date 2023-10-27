package user_primary_adapter

import (
	user_service "go-hexagonal-user-management/core/services/user"
	secondary_port "go-hexagonal-user-management/secondary/port"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FindAll(u user_service.UserService, a secondary_port.UserRepository) *echo.Echo {
	u.Echo.GET("/user", func(c echo.Context) error {
		users, err := u.FindAll(a)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		var responses []UserResponse
		for _, user := range users {
			responses = append(responses, UserResponse{
				Id:       user.Id,
				Username: user.Username,
				Email:    user.Email,
			})
		}
		return c.JSON(http.StatusOK, responses)
	})
	return u.Echo
}
