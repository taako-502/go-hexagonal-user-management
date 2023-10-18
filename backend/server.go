package main

import (
	"net/http"

	"go-sample-api/primary/adapter/user"

	"github.com/labstack/echo/v4"
)

func main() {
		// Echo API(https://echo.labstack.com/)
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.POST("/users", func(c echo.Context) error {
        return user.Create(c)
    })
    e.Logger.Fatal(e.Start(":1323"))
}
