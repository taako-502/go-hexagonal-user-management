package main

import (
	"fmt"
	user_service "go-hexagonal-user-management/core/services/user"
	user_primary_adapter "go-hexagonal-user-management/primary/adapter/user"
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if err := loadEnv(); err != nil {
		log.Fatal(err)
	}

	db := dbInit()

	// DI
	userService := user_service.UserService{}
	userRepo := user_secondary_adapter.NewUserSecondaryAdapter(db)
	pa := user_primary_adapter.NewUserPrimaryAdapter(validator.New())

	// net/http 用のルータ作成
	mux := http.NewServeMux()
	mux.Handle("GET /users", pa.FindAll(userService, userRepo))
	mux.Handle("POST /user", pa.Create(userService, userRepo))
	mux.Handle("PUT /user/{id}", pa.Update(userService, userRepo))
	mux.Handle("DELETE /user/{id}", pa.Delete(userService, userRepo))

	// CORS ミドルウェアを挟む
	handler := corsMiddleware(mux)

	// サーバー起動
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("読み込みできませんでした: %v", err)
	}
	return nil
}

func dbInit() *gorm.DB {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// CORS 設定用のミドルウェア
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3333")
		w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
