package main

import (
	"fmt"
	user_service "go-sample-api/application/services/user"
	"go-sample-api/config"
	user_primary_adapter "go-sample-api/primary/adapter/user"
	user_secondary_adapter "go-sample-api/secondary/adapter/user"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



func main() {
    // 環境変数の読み込み
    loadEnv()
    // Echo API（https://echo.labstack.com/）
    e := echo.New()
		// validation（https://echo.labstack.com/docs/request#validate-data）
		e.Validator = &config.CustomValidator{Validator: validator.New()}
		// CORSの設定（https://echo.labstack.com/docs/middleware/cors）
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:3333"},
			AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
			AllowCredentials: true,
		}))
		// ルーティング（https://echo.labstack.com/docs/routing）
		// TODO: ルーティングの設定を別ファイルに移動する
    e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
    })
		db := dbInit()
		userSecvice := user_service.UserService{Echo: e}
		userSecondaryAdapter := user_secondary_adapter.NewUserSecondaryAdapter(db)
		e = user_primary_adapter.FindAll(userSecvice, userSecondaryAdapter)
		e = user_primary_adapter.Create(userSecvice, userSecondaryAdapter)
		e = user_primary_adapter.Delete(userSecvice, userSecondaryAdapter)
		// サーバー起動
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
