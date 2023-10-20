package user_primary_adapter

import (
	"errors"
	"go-sample-api/application/domain"
	user_service "go-sample-api/application/services/user"
	user_secondary_adapter "go-sample-api/secondary/adapter/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type myDB struct {
	*gorm.DB
}

func NewMyDB(db *gorm.DB) *myDB {
	return &myDB{DB: db}
}

func (db *myDB)Create(c echo.Context) error {
	user := new(domain.User)
	c.Bind(user)
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