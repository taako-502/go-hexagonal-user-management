package user_primary_adapter

import (
	user_service "go-sample-api/application/services/user"
	user_secondary_adapter "go-sample-api/secondary/adapter/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (db *myDB) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, "ID must be an integer")
	}
	a := user_secondary_adapter.NewUserSecondaryAdapter(db.DB)
	if err := user_service.Delete(a , id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "OK")
}