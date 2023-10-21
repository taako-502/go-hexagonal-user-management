package user_primary_adapter

import (
	user_service "go-sample-api/application/services/user"
	user_secondary_adapter "go-sample-api/secondary/adapter/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (db *myDB) FindAll(c echo.Context) error {
	a := user_secondary_adapter.NewUserSecondaryAdapter(db.DB)
	users, err := user_service.FindAll(a)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}