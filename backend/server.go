package main

import (
	"fmt"
	"go-hexagonal-user-management/config"
	user_service "go-hexagonal-user-management/core/services/user"
	user_primary_adapter "go-hexagonal-user-management/primary/adapter/user"
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
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
		AllowOrigins:     []string{"http://localhost:3333"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
	// ルーティング（https://echo.labstack.com/docs/routing）
	// TODO: ルーティングの設定を別ファイルに移動する
	db := dbInit()

	// net/http 用のルータ作成
	mux := http.NewServeMux()
	mux.Handle("GET /users", user_primary_adapter.FindAll(userService, userRepo))
	mux.Handle("POST /user", user_primary_adapter.Create(userService, userRepo))
	mux.Handle("PUT /user/:id", user_primary_adapter.Update(userService, userRepo))
	mux.Handle("DELETE /user/:id", user_primary_adapter.Delete(userService, userRepo))
	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Errorf("読み込み出来ませんでした: %v\n", err)
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
