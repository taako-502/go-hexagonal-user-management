package main

import (
	"fmt"
	user_primary_adapter "go-sample-api/primary/adapter/user"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
    // 環境変数の読み込み
    loadEnv()
    // Echo API(https://echo.labstack.com/)
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    db := dbInit()
    mydb := user_primary_adapter.NewMyDB(db)
    e.POST("/user", mydb.Create)
    e.Logger.Fatal(e.Start(":1323"))
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	} 
}

func dbInit() *gorm.DB {
	dsn := os.Getenv("DSN")
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}


