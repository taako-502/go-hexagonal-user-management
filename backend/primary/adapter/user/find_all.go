package user_primary_adapter

import (
	user_service "go-sample-api/application/services/user"
	user_primary_port "go-sample-api/primary/port/user"
	user_secondary_adapter "go-sample-api/secondary/adapter/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FindAll(u user_service.UserService) *echo.Echo {
	u.Echo.GET("/user", func(c echo.Context) error {
		a := user_secondary_adapter.NewUserSecondaryAdapter(u.DB)
		users, err := u.FindAll(a)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	
		var responses []user_primary_port.UserResponse
		for _, user := range users {
			responses = append(responses, user_primary_port.UserResponse{
				Id: user.Id,
				Username: user.Username,
				Email: user.Email,
			})
		}
		return c.JSON(http.StatusOK, responses)
	})
	return u.Echo
}