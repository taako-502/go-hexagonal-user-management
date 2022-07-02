package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Hello() {
	//http://localhost:8080/helloで"Hello World!"を返却する
	log.Println("start server...")
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	log.Fatal(r.Run())
}
