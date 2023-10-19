package main

import (
	user_primary_adapter "go-sample-api/primary/adapter/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
    // Echo API(https://echo.labstack.com/)
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.POST("/user", user_primary_adapter.Create)
    e.Logger.Fatal(e.Start(":1323"))
}
